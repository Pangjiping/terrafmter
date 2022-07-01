package format

import (
	"fmt"
	"github.com/Pangjiping/terrafmtter/util"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/Pangjiping/terrafmtter/net"
	"github.com/Pangjiping/terrafmtter/terraform"
)

type Resource struct {
	Region  string
	Names   []string
	version string
	// first key is resource name
	// second key is fields
	Fields map[string]map[string]interface{}
}

func NewResource(version string, rs []string) (*Resource, error) {
	var (
		region string
		err    error
	)

	r, valid := validateResource(rs)
	if !valid {
		return nil, fmt.Errorf("invalid resource type: %s", r)
	}

	if region, err = getRegion(); err != nil {
		return nil, err
	}
	names := make([]string, 0)
	for _, r := range rs {
		names = append(names, terraform.ResourceMap[r])
	}
	return &Resource{
		Names:   names,
		version: version,
		Region:  region,
		Fields:  make(map[string]map[string]interface{}),
	}, nil
}

func (r *Resource) Format() error {
	if err := r.getHtmlDocText(); err != nil {
		return err
	}
	if err := r.scan(); err != nil {
		return err
	}
	return nil
}

func (r *Resource) Cleanup() {
	dir, _ := ioutil.ReadDir(util.FILE_LOC_PREFIX)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"", d.Name()}...))
	}
}

// getHtmlCodeText returns parsed code text.
// Deprecated
func (r *Resource) getHtmlCodeText() error {
	for _, re := range r.Names {
		_, err := net.GetCodeFromGithub(r.version, re)
		if err != nil {
			return err
		}
	}
	return nil
}

// getHtmlDocText returns parse markdown doc text.
// enhancement goroutine
func (r *Resource) getHtmlDocText() error {
	total := len(r.Names)
	errChan := make(chan error, total) // accept errors
	respChan := make(chan struct{})
	defer close(errChan)
	defer close(respChan)
	for _, re := range r.Names {
		go func(reso string) {
			if err := net.GetDocFromGithubV2(r.version, reso, true); err != nil {
				errChan <- err
				return
			}
			respChan <- struct{}{}
		}(re)
	}

	// error handle
	for {
		if total == 0 {
			break
		}
		select {
		case err := <-errChan:
			return err
		case <-respChan:
			total--
		}
	}

	return nil
}

func (r *Resource) initRegex() error {
	return nil
}

// todo: goroutine
// scan scans markdown files and parse them into fields.
func (r *Resource) scan() error {
	for _, re := range r.Names {
		resourceName := strings.Join([]string{"alicloud_", re}, "")
		r.Fields[resourceName] = make(map[string]interface{})
		res, err := util.ParseResource(re)
		if err != nil {
			return err
		}
		r.Fields[resourceName] = res.Arguments
	}
	return nil
}

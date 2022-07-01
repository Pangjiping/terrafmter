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
	names   []string
	version string
	// first key is resource name
	// second key is fields
	Fields map[string]map[string]interface{}
}

func NewResource(version string, rs []string) (Formatter, error) {
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
		names:   names,
		version: version,
		Region:  region,
		Fields:  make(map[string]map[string]interface{}),
	}, nil
}

func (r *Resource) Format() error {
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
	for _, re := range r.names {
		_, err := net.GetCodeFromGithub(r.version, re)
		if err != nil {
			return err
		}
	}
	return nil
}

// getHtmlDocText returns parse markdown doc text.
func (r *Resource) getHtmlDocText() error {
	for _, re := range r.names {
		if err := net.GetDocFromGithubV2(r.version, re, true); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resource) initRegex() error {
	return nil
}

// todo: more simple?
func (r *Resource) scan() error {
	for _, re := range r.names {
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

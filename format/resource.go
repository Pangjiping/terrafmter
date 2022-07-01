package format

import (
	"fmt"
	"github.com/Pangjiping/terrafmtter/util"
	"io/ioutil"
	"os"
	"path"

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
	return &Resource{
		names:   rs,
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
		file := terraform.ResourceMap[re]
		_, err := net.GetCodeFromGithub(r.version, file)
		if err != nil {
			return err
		}
	}
	return nil
}

// getHtmlDocText returns parse markdown doc text.
func (r *Resource) getHtmlDocText() error {
	for _, re := range r.names {
		file := terraform.ResourceMap[re]
		if err := net.GetDocFromGithubV2(r.version, file, true); err != nil {
			return err
		}
	}
	return nil
}

func (r *Resource) initRegex() error {
	return nil
}

// todo: scan .md and parse field to r.Fields
func (r *Resource) scan() error {
	return nil
}

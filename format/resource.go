package format

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/Pangjiping/terrafmtter/net"
	"github.com/Pangjiping/terrafmtter/terraform"
)

type Resource struct {
	Region  string
	name    string
	version string
	Fields  map[string]Field
}

func NewResource(version, r string) (Formatter, error) {
	var (
		region string
		err    error
	)

	if valid := validateResource(r); !valid {
		return nil, fmt.Errorf("invalid resource type: %s", r)
	}

	if region, err = getRegion(); err != nil {
		return nil, err
	}
	return &Resource{
		name:    r,
		version: version,
		Region:  region,
		Fields:  make(map[string]Field),
	}, nil
}

func (r *Resource) Format() error {
	return nil
}

func (r *Resource) Cleanup() {
	dir, _ := ioutil.ReadDir(net.FILE_LOC_PREFIX)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"", d.Name()}...))
	}
}

// getHtmlCodeText returns parsed code text.
// Deprecated
func (r *Resource) getHtmlCodeText() (string, error) {
	file := terraform.ResourceMap[r.name]
	return net.GetCodeFromGithub(r.version, file)
}

// getHtmlDocText returns parse markdown doc text.
func (r *Resource) getHtmlDocText() error {
	file := terraform.ResourceMap[r.name]
	if err := net.GetDocFromGithubV2(r.version, file, true); err != nil {
		return err
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

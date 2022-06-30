package format

import (
	"fmt"
	"github.com/Pangjiping/terrafmter/net"
	"github.com/Pangjiping/terrafmter/terraform"
)

type Resource struct {
	AccessKey string
	SecretKey string
	Region    string
	name      string
	htmlCache string
	version   string
	Fields    map[string]Field
}

func NewResource(version, r string) (Formatter, error) {
	var (
		accKey string
		secKey string
		region string
		err    error
	)

	if valid := validateResource(r); !valid {
		return nil, fmt.Errorf("invalid resource type: %s", r)
	}

	if accKey, err = getAccessKey(); err != nil {
		return nil, err
	}
	if secKey, err = getSecretKey(); err != nil {
		return nil, err
	}
	if region, err = getRegion(); err != nil {
		return nil, err
	}
	return &Resource{
		name:      r,
		version:   version,
		AccessKey: accKey,
		SecretKey: secKey,
		Region:    region,
		Fields:    make(map[string]Field),
	}, nil
}

func (r *Resource) Format() error {
	return nil
}

func (r *Resource) Cleanup() {

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
	s, err := net.GetDocFromGithub(r.version, file, true)
	if err != nil {
		return err
	}
	r.htmlCache = s
	return nil
}

func (r *Resource) initRegex() error {
	return nil
}

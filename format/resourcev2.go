package format

import (
	"fmt"
	"github.com/Pangjiping/terrafmtter/net"
	"github.com/Pangjiping/terrafmtter/util"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// source and data
const (
	DATA       = "data"
	RESOURCE   = "resource"
	url_format = "https://registry.terraform.io/providers/aliyun/alicloud/%s/docs/%s/%s"
)

type SchemaMapping struct {
	Type    string  `json:"type"`    // like data and resource
	Name    string  `json:"name"`    // like cs_kubernetes, cs_managed_kubernetes
	Version string  `json:"version"` // like 1.173.0
	Url     string  `json:"url"`     // doc url from hashicorp
	Fields  []Field `json:"fields"`  // schema
}

type Field struct {
	Name        string   `json:"name"`
	Description []string `json:"description"` // description need to split by .
	Optional    bool     `json:"optional"`
	Required    bool     `json:"required"`
	ForceNew    bool     `json:"force_new"`
	Detail      string   `json:"detail"` // details url
}

// NewSchemaMapping get SchemaMapping
// name like kubernetes
func NewSchemaMapping(version string, typ string, name string) *SchemaMapping {
	return &SchemaMapping{
		Type:    typ,
		Name:    name,
		Version: version,
		Fields:  make([]Field, 0),
	}
}

func (sm *SchemaMapping) Format() error {
	defer sm.cleanup()
	sm.getUrl()
	if err := sm.getDocs(); err != nil {
		return err
	}
	if err := sm.scan(); err != nil {
		return err
	}
	return nil
}

func (sm *SchemaMapping) cleanup() {
	dir, _ := ioutil.ReadDir(util.FILE_LOC_PREFIX)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{util.FILE_LOC_PREFIX, d.Name()}...))
	}
}

func (sm *SchemaMapping) scan() error {
	res, err := parseResource(sm.Name)
	if err != nil {
		return err
	}
	if err = sm.convertParsed2Fields(res); err != nil {
		return err
	}
	return nil
}

// getDocs fetch doc from github.
func (sm *SchemaMapping) getDocs() error {
	switch sm.Type {
	case DATA:
		if err := net.GetDocFromGithubV2(sm.Version, sm.Name, false); err != nil {
			return err
		}
	case RESOURCE:
		if err := net.GetDocFromGithubV2(sm.Version, sm.Name, true); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid resource type: %s", sm.Type)
	}
	return nil
}

func (sm *SchemaMapping) convertParsed2Fields(prev *parsed) error {
	field := Field{}
	for _, v1 := range prev.arguments {
		field.Name = v1[NAME]
		field.Detail = strings.Join([]string{sm.Url, field.Name}, "#")

		// handle property
		for k2, v2 := range v1 {
			switch k2 {
			case OPTIONAL:
				if v2 == "true" {
					field.Optional = true
				}
			case FORCENEW:
				if v2 == "true" {
					field.ForceNew = true
				}
			case REQUIRED:
				if v2 == "true" {
					field.Required = true
				}
			default:
			}
		}
		sm.Fields = append(sm.Fields, field)
	}
	return nil
}

func (sm *SchemaMapping) getUrl() {
	if sm.Type == DATA {
		sm.Url = fmt.Sprintf(url_format, sm.Version, "data-sources", sm.Name)
	} else {
		sm.Url = fmt.Sprintf(url_format, sm.Version, "resources", sm.Name)
	}
}

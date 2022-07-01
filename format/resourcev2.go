package format

import (
	"fmt"
	"github.com/Pangjiping/terrafmtter/net"
	"strings"
)

// source and data
const (
	DATA     = "data"
	RESOURCE = "resource"
)

type SchemaMapping struct {
	Type    string  // like data and resource
	Name    string  // like cs_kubernetes, cs_managed_kubernetes
	Version string  // like 1.173.0
	Fields  []Field // schema
}

type Field struct {
	Name        string
	Description []string // description need to split by .
	OptOrReq    string
	ForceNew    bool
}

// NewSchemaMapping get SchemaMapping
// name like kubernetes
func NewSchemaMapping(version string, typ string, name string) *SchemaMapping {
	fullName := strings.Join([]string{"cs", name}, "_")
	return &SchemaMapping{
		Type:    typ,
		Name:    fullName,
		Version: version,
		Fields:  make([]Field, 0),
	}
}

func (sm *SchemaMapping) Format() error {
	defer sm.cleanup()
	if err := sm.getDocs(); err != nil {
		return err
	}
	if err := sm.scan(); err != nil {
		return err
	}
	return nil
}

func (sm *SchemaMapping) cleanup() {
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
	for k1, v1 := range prev.arguments {
		field.Name = k1
		for k2, v2 := range v1 {
			switch k2 {
			case OPTIONAL:
				if v2 == "true" {
					field.OptOrReq = OPTIONAL
				}
			case FORCENEW:
				if v2 == "true" {
					field.ForceNew = true
				}
			case REQUIRED:
				if v2 == "true" {
					field.OptOrReq = REQUIRED
				}
			case DESCRIPTION:
				descriptions := strings.Split(v2, ".")
				field.Description = descriptions
			default:
			}
		}
		sm.Fields = append(sm.Fields, field)
	}
	return nil
}

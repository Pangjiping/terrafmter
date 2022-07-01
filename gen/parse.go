package gen

import (
	"fmt"
	"github.com/Pangjiping/terrafmtter/format"
	"github.com/Pangjiping/terrafmtter/util"
	"os"
	"sync"
	"text/template"
)

var (
	mappings []format.SchemaMapping
	rwlock   sync.RWMutex
)

func init() {
	mappings = make([]format.SchemaMapping, 0)
	rwlock = sync.RWMutex{}
}

func Execute(resources, datas []string, fileName, version string) error {

	// valid data and resource
	r, ok := util.ValidateResource(resources)
	if !ok {
		return fmt.Errorf("Invalid resource type: %s", r)
	}
	d, ok := util.ValidateDataSource(datas)
	if !ok {
		return fmt.Errorf("Invalid data source type: %s", d)
	}

	// schemaMapping format
	if err := formatMapping(resources, datas, version); err != nil {
		return err
	}

	// write to target file
	if err := renderFile(); err != nil {
		return err
	}

	return nil
}

// formatMapping parse .md file to SchemaMapping
// goroutine
func formatMapping(resources, datas []string, version string) error {
	total := len(resources) + len(datas)
	errChan := make(chan error, total)
	respChan := make(chan struct{})
	defer close(errChan)
	defer close(respChan)

	// datasource
	for _, ds := range datas {
		go func(data, version string) {
			mapping := format.NewSchemaMapping(version, "data", data)
			if err := mapping.Format(); err != nil {
				errChan <- err
				return
			}
			rwlock.Lock()
			mappings = append(mappings, *mapping)
			rwlock.Unlock()
			respChan <- struct{}{}
		}(ds, version)
	}

	// resource
	for _, re := range resources {
		go func(resource, version string) {
			mapping := format.NewSchemaMapping(version, "resource", resource)
			if err := mapping.Format(); err != nil {
				errChan <- err
				return
			}
			rwlock.Lock()
			mappings = append(mappings, *mapping)
			rwlock.Unlock()
			respChan <- struct{}{}
		}(re, version)
	}

	// error handler
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

func renderFile() error {

	// the name inc is what the function will be called in the template text.
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}

	tpl, err := template.New("").Funcs(funcMap).Parse(terraformBaseTemplate)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("main.tf", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = tpl.Execute(f, mappings); err != nil {
		return err
	}
	return nil
}

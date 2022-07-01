package gen

import (
	"embed"
	"fmt"
	"github.com/Pangjiping/terrafmtter/format"
	"github.com/Pangjiping/terrafmtter/util"
	"sync"
)

//go:embed templates/*
var tmpl embed.FS

var (
	mappings []*format.SchemaMapping
	rwlock   sync.RWMutex
)

func init() {
	mappings = make([]*format.SchemaMapping, 0)
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

	// 渲染模版文件

	// 写文件

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
			mappings = append(mappings, mapping)
			rwlock.Unlock()
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
			mappings = append(mappings, mapping)
			rwlock.Unlock()
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

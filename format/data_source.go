package format

import (
	"fmt"
	"github.com/Pangjiping/terrafmtter/net"
	"github.com/Pangjiping/terrafmtter/terraform"
	"github.com/Pangjiping/terrafmtter/util"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type DataSource struct {
	Region  string
	Names   []string
	version string
	Fields  map[string]map[string]interface{}
}

func NewDataSource(version string, ds []string) (Formatter, error) {
	var (
		region string
		err    error
	)
	d, valid := validateDataSource(ds)
	if !valid {
		return nil, fmt.Errorf("invalid data source type: %s", d)
	}
	if region, err = getRegion(); err != nil {
		return nil, err
	}
	names := make([]string, 0)
	for _, d := range ds {
		names = append(names, terraform.DataSourceMap[d])
	}
	return &DataSource{
		Names:   names,
		version: version,
		Region:  region,
		Fields:  make(map[string]map[string]interface{}),
	}, nil

}

func (d *DataSource) Format() error {
	if err := d.getHtmlDocText(); err != nil {
		return err
	}
	if err := d.scan(); err != nil {
		return err
	}
	return nil
}

func (d *DataSource) Cleanup() {
	dir, _ := ioutil.ReadDir(util.FILE_LOC_PREFIX)
	for _, file := range dir {
		os.RemoveAll(path.Join([]string{"", file.Name()}...))
	}
}

func (d *DataSource) getHtmlDocText() error {
	total := len(d.Names)
	errChan := make(chan error, total) // accept errors
	respChan := make(chan struct{})
	defer close(errChan)
	defer close(respChan)
	for _, da := range d.Names {
		go func(data string) {
			if err := net.GetDocFromGithubV2(d.version, data, false); err != nil {
				errChan <- err
				return
			}
			respChan <- struct{}{}
		}(da)
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

func (d *DataSource) scan() error {
	for _, da := range d.Names {
		dataSourceName := strings.Join([]string{"alicloud_", da}, "")
		d.Fields[dataSourceName] = make(map[string]interface{})
		res, err := util.ParseResource(da)
		if err != nil {
			return err
		}
		d.Fields[dataSourceName] = res.Arguments
	}
	return nil
}

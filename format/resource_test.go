package format

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	region  = "cn-hangzhou"
	version = "1.173.0"
)

var (
	resources = []string{
		// resource
		"cs_kubernetes",
		"cs_managed_kubernetes",
		"cs_kubernetes_node_pool",
	}

	rs = []string{
		"kubernetes",
		"managed_kubernetes",
		"kubernetes_node_pool",
	}
)

func mockResource(resource []string) *Resource {
	return &Resource{
		names:   resource,
		Region:  region,
		version: version,
		Fields:  make(map[string]map[string]interface{}),
	}
}

// to download .md file and parse it to fields
func TestResourceFormat(t *testing.T) {
	resClient, err := NewResource(version, rs)
	assert.Nil(t, err)
	assert.NotNil(t, resClient)
	err = resClient.Format()
	assert.Nil(t, err)
	assert.NotNil(t, resClient.Fields)
	fmt.Println(resClient.Fields)
}

func TestGetDocFromWeb(t *testing.T) {
	resClient, err := NewResource(version, rs)
	assert.Nil(t, err)
	assert.NotNil(t, resClient)
	err = resClient.getHtmlDocText()
	assert.Nil(t, err)
}

func TestScan(t *testing.T) {
	resClient := mockResource(resources)
	err := resClient.scan()
	assert.Nil(t, err)
	fmt.Println(resClient.Fields)
}

func TestCleanup(t *testing.T) {
	resClient := mockResource(resources)
	assert.NotNil(t, resClient)
	resClient.Cleanup()
}

//func TestGetHtmlDocText(t *testing.T) {
//	for k := range terraform.ResourceMap {
//		reso := mockResource(k)
//		assert.NotNil(t, reso)
//		err := reso.getHtmlDocText()
//		assert.Nil(t, err)
//		assert.NotNil(t, reso.htmlCache)
//	}
//}

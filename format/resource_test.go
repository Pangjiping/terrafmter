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
)

func mockResource(resource []string) *Resource {
	return &Resource{
		names:   resource,
		Region:  region,
		version: version,
		Fields:  make(map[string]map[string]interface{}),
	}
}

func TestScan(t *testing.T) {
	resClient := mockResource(resources)
	err := resClient.scan()
	assert.Nil(t, err)
	fmt.Println(resClient.Fields)
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

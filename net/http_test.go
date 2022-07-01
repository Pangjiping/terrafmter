package net

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	version = "1.173.0"
	file    = "data_source_alicloud_cs_kubernetes_version.go"
	doc     = "cs_kubernetes_version"
	datas   = []string{
		// data source
		"cs_kubernetes_version",
		"cs_managed_kubernetes_clusters",
		"cs_kubernetes_addon_metadata",
	}
	resources = []string{
		// resource
		"cs_kubernetes",
		"cs_managed_kubernetes",
		"cs_kubernetes_node_pool",
	}
)

func TestGetHttpAddr(t *testing.T) {
	t.SkipNow()
	addr := getHttpAddr(version, file)
	assert.NotNil(t, addr)
	assert.Equal(t, `https://github.com/aliyun/terraform-provider-alicloud/blob/v1.173.0/alicloud/data_source_alicloud_cs_kubernetes_version.go`, addr)
}

// todo: add test case
func TestGetCodeFromGithub(t *testing.T) {
	t.SkipNow()
	s, err := GetCodeFromGithub(version, file)
	assert.Nil(t, err)
	assert.NotNil(t, s)
}

// todo: add test case
func TestGetDocFromGithub(t *testing.T) {
	t.SkipNow()
	s, err := GetDocFromGithub(version, doc, false)
	assert.Nil(t, err)
	assert.NotNil(t, s)
}

// get more datasource and resource
// todo: more error testcase
func TestGetDocFromGithubV2(t *testing.T) {
	for _, tar := range datas {
		err := GetDocFromGithubV2(version, tar, false)
		assert.Nil(t, err)
	}
	for _, tar := range resources {
		err := GetDocFromGithubV2(version, tar, true)
		assert.Nil(t, err)
	}
}

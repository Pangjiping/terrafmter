package net

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	version = "1.173.0"
	file    = "data_source_alicloud_cs_kubernetes_version.go"
	doc     = "cs_kubernetes_version"
)

func TestGetHttpAddr(t *testing.T) {
	t.SkipNow()
	addr := getHttpAddr(version, file)
	assert.NotNil(t, addr)
	assert.Equal(t, `https://github.com/aliyun/terraform-provider-alicloud/blob/v1.173.0/alicloud/data_source_alicloud_cs_kubernetes_version.go`, addr)
}

// todo: add test case
func TestGetCodeFromGithub(t *testing.T) {
	s, err := GetCodeFromGithub(version, file)
	assert.Nil(t, err)
	assert.NotNil(t, s)
}

// todo: add test case
func TestGetDocFromGithub(t *testing.T) {
	s, err := GetDocFromGithub(version, doc, false)
	assert.Nil(t, err)
	assert.NotNil(t, s)
}

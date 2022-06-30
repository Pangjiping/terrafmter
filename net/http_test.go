package net

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHttpAddr(t *testing.T) {
	t.SkipNow()
	version := "1.173.0"
	file := "data_source_alicloud_cs_kubernetes_version.go"
	addr := getHttpAddr(version, file)
	assert.NotNil(t, addr)
	assert.Equal(t, `https://github.com/aliyun/terraform-provider-alicloud/blob/v1.173.0/alicloud/data_source_alicloud_cs_kubernetes_version.go`, addr)
}

func TestGetCodeFromGithub(t *testing.T) {
	version := "1.173.0"
	file := "data_source_alicloud_cs_kubernetes_version.go"
	addr := getHttpAddr(version, file)
	assert.NotNil(t, addr)
	assert.Equal(t, `https://github.com/aliyun/terraform-provider-alicloud/blob/v1.173.0/alicloud/data_source_alicloud_cs_kubernetes_version.go`, addr)

	s, err := GetCodeFromGithub(version, file)
	assert.Nil(t, err)
	assert.NotNil(t, s, s)
}

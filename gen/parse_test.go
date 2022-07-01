package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	resources = []string{
		"cs_managed_kubernetes",
	}
	datas = []string{
		"cs_kubernetes_version",
	}
	version  = "1.173.0"
	fileName = "main.tf"
)

func TestExecute(t *testing.T) {
	err := Execute(resources, datas, fileName, version)
	assert.Nil(t, err)
}

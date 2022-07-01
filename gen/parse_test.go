package gen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	resources = []string{
		"managed_kubernetes",
	}
	datas = []string{
		"kubernetes_version",
	}
	version  = "1.173.0"
	fileName = "main.tf"
)

func TestExecute(t *testing.T) {
	err := Execute(resources, datas, fileName, version)
	assert.Nil(t, err)
}

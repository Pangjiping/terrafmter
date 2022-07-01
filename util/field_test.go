package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo: more test case for resource and datasource
func TestParseResource(t *testing.T) {
	resourceName := "cs_kubernetes_version"
	res, err := parseResource(resourceName)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	fmt.Println(*res)
}

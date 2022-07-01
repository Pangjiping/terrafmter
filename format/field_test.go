package format

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo: more test case for resource and datasource
func TestParseResource(t *testing.T) {
	testcase := map[string]interface{}{
		"Error": false,
		"Resources": []string{
			// data source
			"cs_kubernetes_version",
			"cs_managed_kubernetes_clusters",
			"cs_kubernetes_addon_metadata",

			// resource
			"cs_kubernetes",
			"cs_managed_kubernetes",
			"cs_kubernetes_node_pool",
		},
		"Wanted": true,
	}
	for _, r := range testcase["Resources"].([]string) {
		res, err := parseResource(r)
		assert.Equal(t, []interface{}{
			err == nil,
			res != nil,
		}, []interface{}{
			testcase["Wanted"].(bool),
			testcase["Wanted"].(bool),
		})
		fmt.Println(*res)
	}
}

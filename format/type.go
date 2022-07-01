package format

import "github.com/Pangjiping/terrafmtter/terraform"

const (
	EMPTY = "nil"

	// accessKey and secretKey
	ACCESS_KEY     = "ALICLOUD_ACCESS_KEY"
	SECRET_KEY     = "ALICLOUD_SECRET_KEY"
	REGION         = "ALICLOUD_REGION"
	DEFAULT_REGION = "cn-hangzhou"
)

// validateType 检查resource和data不能同时设置，也不能都不设置
func validateType(r, d string) bool {
	if r == EMPTY && d == EMPTY {
		return false
	} else if r != EMPTY && d != EMPTY {
		return false
	}
	return true
}

func validateResource(rs []string) (string, bool) {
	for _, r := range rs {
		if _, ok := terraform.ResourceMap[r]; !ok {
			return r, false
		}
	}
	return "", true
}

func validateDataSource(ds []string) (string, bool) {
	for _, d := range ds {
		if _, ok := terraform.DataSourceMap[d]; !ok {
			return d, false
		}
	}
	return "", true
}

type response struct {
	finished bool
	err      error
}

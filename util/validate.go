package util

import "github.com/Pangjiping/terrafmtter/terraform"

func ValidateResource(rs []string) (string, bool) {
	for _, r := range rs {
		if _, ok := terraform.ResourceMap[r]; !ok {
			return r, false
		}
	}
	return "", true
}

func ValidateDataSource(ds []string) (string, bool) {
	for _, d := range ds {
		if _, ok := terraform.DataSourceMap[d]; !ok {
			return d, false
		}
	}
	return "", true
}

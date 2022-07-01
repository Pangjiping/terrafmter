package format

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

type response struct {
	finished bool
	err      error
}

package format

var (
	region  = "cn-hangzhou"
	version = "1.173.0"
)

func mockResource(resource []string) *Resource {
	return &Resource{
		names:   resource,
		Region:  region,
		version: version,
		Fields:  make(map[string]map[string]interface{}),
	}
}

//func TestGetHtmlDocText(t *testing.T) {
//	for k := range terraform.ResourceMap {
//		reso := mockResource(k)
//		assert.NotNil(t, reso)
//		err := reso.getHtmlDocText()
//		assert.Nil(t, err)
//		assert.NotNil(t, reso.htmlCache)
//	}
//}

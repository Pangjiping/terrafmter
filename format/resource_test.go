package format

import (
	"github.com/Pangjiping/terrafmter/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	accKey  = "hauifhakjsnfkaj"
	secKey  = "fsuGAfakdnaGYmsnmxn"
	region  = "cn-hangzhou"
	version = "1.173.0"
)

func mockResource(resource string) *Resource {
	return &Resource{
		name:      resource,
		Region:    region,
		AccessKey: accKey,
		SecretKey: secKey,
		version:   version,
		Fields:    make(map[string]Field),
	}
}

func TestGetHtmlDocText(t *testing.T) {
	for k := range terraform.ResourceMap {
		reso := mockResource(k)
		assert.NotNil(t, reso)
		err := reso.getHtmlDocText()
		assert.Nil(t, err)
		assert.NotNil(t, reso.htmlCache)
	}
}

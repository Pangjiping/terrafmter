package net

import (
	"fmt"
	"github.com/Pangjiping/terrafmter/util"
	"io/ioutil"
	"net/http"
)

func GetCodeFromGithub(version, file string) (string, error) {
	resp, err := http.Get(getHttpAddr(version, file))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	s := util.Bytes2String(body)
	err = util.WriteToFile("test.go", s)
	return s, err
}

func getHttpAddr(version, file string) string {
	return fmt.Sprintf(`https://github.com/aliyun/terraform-provider-alicloud/blob/v%s/alicloud/%s`, version, file)
}

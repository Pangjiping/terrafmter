package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func WriteToFile(filePath string, data interface{}) error {
	var out string
	switch data.(type) {
	case string:
		out = data.(string)
		break
	case nil:
		return nil
	default:
		bs, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return fmt.Errorf("MarshalIndent data %#v got an error: %#v", data, err)
		}
		out = Bytes2String(bs)
	}

	if _, err := os.Stat(filePath); err == nil {
		if err = os.Remove(filePath); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(filePath, String2Bytes(out), 422)
}

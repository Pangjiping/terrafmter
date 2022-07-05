package gen

import (
	"encoding/json"

	"github.com/Pangjiping/terrafmtter/util"
)

// mapping2Json for test.
func mappingToJson() error {
	b, err := json.Marshal(mappings)
	if err != nil {
		return err
	}
	err = util.WriteToFile("parse.json", util.Bytes2String(b))
	return err
}

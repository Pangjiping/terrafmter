package ack

// attach到一个ack资源可能需要的所有resource和datasource

// attachMap key: resource or data
//           value: name
var attachMap map[string]string

func init() {
	attachMap = map[string]string{}
}

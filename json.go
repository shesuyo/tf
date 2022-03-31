package tf

import (
	"encoding/json"
)

// JSONStringify JSON字符串化
func JSONStringify(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	return BytesString(bs)
}

// JSONBytes 将对象转化成JSON字符串化的[]byte类型
func JSONBytes(obj interface{}) []byte {
	bs, _ := JSONMarshal(obj)
	return bs
}

// JSONUnmarshal json.Unmarshal
func JSONUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// JSONMarshal JSONMarshal
func JSONMarshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

// JSONMapString bytes to map[string]interface{}
func JSONMapString(data []byte) map[string]interface{} {
	m := make(map[string]interface{}, 0)
	json.Unmarshal(data, &m)
	return m
}

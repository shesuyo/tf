package wtf

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
	bs, _ := json.Marshal(obj)
	return bs
}

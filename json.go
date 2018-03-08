package wtf

import (
	"encoding/json"
)

func JSONStringify(obj interface{}) string {
	bs, _ := json.Marshal(obj)
	return BytesString(bs)
}

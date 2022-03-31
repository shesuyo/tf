package tf

import (
	"reflect"
)

// Struct2SI 把结构体转换成map[string]interface{}
func Struct2SI(s interface{}) map[string]interface{} {
	v := reflect.ValueOf(s)
	v = reflect.Indirect(v)
	t := v.Type()
	m := map[string]interface{}{}
	for i, num := 0, v.NumField(); i < num; i++ {
		tag := t.Field(i).Tag
		field := tag.Get("json")
		if field == "" {
			// field = t.Field(i).Name
			continue
		}
		m[field] = v.Field(i).Interface()
	}
	return m
}

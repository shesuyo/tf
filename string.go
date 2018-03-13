package wtf

import (
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

// String 将传入的值转换成字符串
func String(v interface{}) string {
	var s string
	switch v := v.(type) {
	case int:
		s = strconv.Itoa(v)
	case int64:
		s = strconv.Itoa(int(v))
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

// StringsContain a集合中是否包含b元素
func StringsContain(a []string, b string) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

// StringBytes 将string转成[]byte
func StringBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// BytesString 将[]byte转成string
func BytesString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// StringNoSpace 返回一个没有空格的字符串
func StringNoSpace(s string) string {
	return strings.Replace(s, " ", "", -1)
}

// ***************************************************
// ********************strings************************
// ***************************************************

// Strings2Interfaces 将[]string转换成[]interface{}
func Strings2Interfaces(ss []string) []interface{} {
	is := make([]interface{}, len(ss))
	for k, v := range ss {
		is[k] = v
	}
	return is
}

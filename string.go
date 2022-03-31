package tf

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// String 将传入的值转换成字符串
func String(v interface{}) string {
	if v == nil {
		return ""
	}
	var s string
	switch v := v.(type) {
	case int:
		s = strconv.Itoa(v)
	case int64:
		s = strconv.Itoa(int(v))
	case string:
		return v
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

var stringRandomBytes = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}

// StringRandom 返回随机数随机数的字符串构成为0-9a-zA-Z
func StringRandom(n int) string {
	sb := strings.Builder{}
	for i := 0; i < n; i++ {
		num := gRand.Intn(61)
		sb.WriteByte(stringRandomBytes[num])
	}
	return sb.String()
}

var gRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringRandomNumber 返回随机的digits位数字
func StringRandomNumber(digits int) string {
	sb := strings.Builder{}
	for i := 0; i < digits; i++ {
		num := gRand.Int63() % 10
		sb.WriteByte(stringRandomBytes[num])
	}
	return sb.String()
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

// StringsContains a中是否有包含b的元素
func StringsContains(a []string, b string) bool {
	for _, v := range a {
		if strings.Contains(v, b) {
			return true
		}
	}
	return false
}

// StringsContainsReverse b中是否有包含a中任意一个元素
func StringsContainsReverse(a []string, b string) bool {
	for _, v := range a {
		if strings.Contains(b, v) {
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

func Interfaces2Strings(is []interface{}) []string {
	ss := make([]string, len(is))
	for k, v := range is {
		ss[k] = String(v)
	}
	return ss
}

// Strings2Ints 将[]string转换成[]int
func Strings2Ints(ss []string) []int {
	is := make([]int, len(ss))
	for k, v := range ss {
		is[k] = Int(v)
	}
	return is
}

// StringsSplit strings.Split proxy 如果s为空，则返回长度为0的空字符串数组。
func StringsSplit(s, sep string) []string {
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, sep)
}

// StringRepeat string repeat
func StringRepeat(s string, t int) string {
	bs := bytes.Buffer{}
	for i := 0; i < t; i++ {
		bs.WriteString(s)
	}
	return bs.String()
}

// StringRepeatSlice StringRepeatSlice
func StringRepeatSlice(s string, t int) []string {
	sli := make([]string, 0, t)
	for i := 0; i < t; i++ {
		sli = append(sli, s)
	}
	return sli
}

// StringsUnique 去重
func StringsUnique(ss []string) []string {
	m := make(map[string]struct{}, len(ss))
	for _, s := range ss {
		m[s] = struct{}{}
	}
	if len(ss) == len(m) {
		return ss
	}
	nss := make([]string, 0, len(m))
	for s := range m {
		nss = append(nss, s)
	}
	return nss
}

// StringAdd 两个字符串数字相加
func StringAdd(a, b string) string {
	numa, _ := strconv.ParseFloat(a, 64)
	numb, _ := strconv.ParseFloat(b, 64)
	sum := numa + numb
	return String(FloatDecimal2(sum))
}

// StringDecimal2 字符串浮点类型取2位小数
func StringDecimal2(val string) string {
	num, _ := strconv.ParseFloat(val, 64)
	num = FloatDecimal2(num)
	return String(num)
}

// StringTmpName 临时文件名
func StringTmpName() string {
	return strconv.Itoa(int(time.Now().UnixNano())) + StringRandom(7)
}

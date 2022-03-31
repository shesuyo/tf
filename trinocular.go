package tf

import "strings"

//S 三目运算符
func S(s bool, a, b interface{}) interface{} {
	if s {
		return a
	}
	return b
}

//SBool 三目运算符
func SBool(s, a, b bool) bool {
	if s {
		return a
	}
	return b
}

//SInt 三目运算符int类型
func SInt(s bool, a, b int) int {
	if s {
		return a
	}
	return b
}

//SString 三目运算符String类型
func SString(s bool, a, b string) string {
	if s {
		return a
	}
	return b
}

//PlugeString 提取map中某个列为slice
func PlugeString(ms []map[string]string, key string) []string {
	out := []string{}
	for _, m := range ms {
		if val, ok := m[key]; ok {
			out = append(out, val)
		}
	}
	return out
}

//SQLInString SQL IN (1,3,4,5,8,4,5)
func SQLInString(args []string) string {
	return "(" + strings.Join(args, ",") + ")"
}

package tf

import (
	"fmt"
	"sort"
	"strconv"
)

// Int 返回这个值对应的int值
func Int(v interface{}) int {
	var i int
	switch v := v.(type) {
	case string:
		i, _ = strconv.Atoi(v)
	// 一个case多个值，就无法确认是什么类型了，就成了interface{}，所以要分开写。
	case int64:
		i = int(v)
	// 不实现除了uint64之后的无符号
	case uint64:
		i = int(v)
	case int:
		i = v
	case int8:
		i = int(v)
	case int16:
		i = int(v)
	case int32:
		i = int(v)
	default:
		i, _ = strconv.Atoi(fmt.Sprintf("%v", v))
	}
	return i
}

// IntAbs intabs
func IntAbs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// IntDivide 返回a/b的值
func IntDivide(a, b int) int {
	if b == 0 {
		return 0
	}
	return a / b
}

// IntPercent 百分比求值，不保留小数。
func IntPercent(a, b int) int {
	if b == 0 {
		return 0
	}
	return int((float64(a) / float64(b)) * 100)
}

// IntPercentStr 百分比求值，不保留小数。
func IntPercentStr(a, b int) string {
	if b == 0 {
		return "0%"
	}
	return fmt.Sprintf("%.0f%%", (float64(a)/float64(b))*100)
}

// IntPercentPoint1 百分比求值，保留一位小数。
func IntPercentPoint1(a, b int) string {
	if b == 0 {
		return "0%"
	}
	return fmt.Sprintf("%.1f%%", (float64(a)/float64(b))*100)
}

// IntPercentPoint2 百分比求值，保留两位小数。
func IntPercentPoint2(a, b int) string {
	if b == 0 {
		return "0%"
	}
	return fmt.Sprintf("%.2f%%", (float64(a)/float64(b))*100)
}

// ***************************************************
// **********************ints*************************
// ***************************************************

// Ints2Interfaces 将[]int转成[]interface{}
func Ints2Interfaces(a []int) []interface{} {
	is := make([]interface{}, len(a))
	for k, v := range a {
		is[k] = v
	}
	return is
}

type intsDesc []int

func (p intsDesc) Len() int           { return len(p) }
func (p intsDesc) Less(i, j int) bool { return p[i] > p[j] }
func (p intsDesc) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// IntsSortDesc []int降序排
func IntsSortDesc(a []int) {
	sort.Sort(intsDesc(a))
}

// IntsContain a集合中是否包含了b
func IntsContain(a []int, b int) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}

// IntsContains a集合与b集合是否相交
func IntsContains(a, b []int) bool {
	m := make(map[int]bool, len(a))
	for _, v := range a {
		m[v] = true
	}
	for _, v := range b {
		if m[v] {
			return true
		}
	}
	return false
}

// IntsContainsAny a集合是否包含b集合中任一元素
func IntsContainsAny(a, b []int) bool {
	m := make(map[int]struct{}, len(a))
	for _, v := range a {
		m[v] = struct{}{}
	}
	for _, v := range b {
		_, ok := m[v]
		if ok {
			return true
		}
	}
	return false
}

// IntsUnion 返回a和b的集合
func IntsUnion(a, b []int) []int {
	want := []int{}
	want = append(want, a...)
	want = append(want, b...)
	return want
}

// IntsDiff 返回a集合不同于b集合的元素
func IntsDiff(a, b []int) []int {
	want := []int{}
	for _, ae := range a {
		have := false
		for _, be := range b {
			if be == ae {
				have = true
				break
			}
		}
		if !have {
			want = append(want, ae)
		}
	}
	return want
}

// IntsIntersect 返回a集合与b集合的交集
func IntsIntersect(a, b []int) []int {
	want := []int{}
	for _, ae := range a {
		have := false
		for _, be := range b {
			if be == ae {
				have = true
				break
			}
		}
		if have {
			want = append(want, ae)
		}
	}
	return want
}

// IntsUnique 除重
func IntsUnique(ids []int) []int {
	tagMap := make(map[int]bool, len(ids))
	ret := make([]int, 0, len(ids))
	for _, id := range ids {
		if !tagMap[id] {
			tagMap[id] = true
			ret = append(ret, id)
		}
	}
	return ret
}

// FloatDecimal2 FloatDecimal2
func FloatDecimal2(f float64) float64 {
	var val float64
	if f > 0 {
		val = float64(int((f+0.005)*100)) / 100.0
	} else {
		val = float64(int((f-0.005)*100)) / 100.0
	}
	return val
}

func Float(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func Float2(s string) float64 {
	return FloatDecimal2(Float(s))
}

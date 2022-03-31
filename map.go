package tf

import (
	"sort"
)

// 默认为StringString
// string s
// int i
// interface{} if
// bool b

// MapCopy 复制map[string]string
func MapCopy(m map[string]string) map[string]string {
	n := make(map[string]string, len(m))
	for k, v := range m {
		n[k] = v
	}
	return n
}

// MapInitVal 初始化值为val
func MapInitVal(m map[string]string, val string) {
	for k := range m {
		m[k] = val
	}
}

// MapMerge 将tail map合并到head map中
func MapMerge(head map[string]string, tail map[string]string) map[string]string {
	for k, v := range tail {
		head[k] = v
	}
	return head
}

// MapString2Interface 将map[string]string转换成map[string]interface{}
func MapString2Interface(s map[string]string) map[string]interface{} {
	i := make(map[string]interface{}, len(s))
	for k, v := range s {
		i[k] = v
	}
	return i
}

// KVOrderDesc 专门用来排序的
type KVOrderDesc struct {
	Key interface{}
	Val int
}

func (kv KVOrderDesc) String() string {
	str, ok := kv.Key.(string)
	if ok {
		return str
	}
	return ""
}

func (kv KVOrderDesc) Int() int {
	str, ok := kv.Key.(int)
	if ok {
		return str
	}
	return 0
}

// KVOrderDescs []KV
type KVOrderDescs []KVOrderDesc

func (p KVOrderDescs) Len() int           { return len(p) }
func (p KVOrderDescs) Less(i, j int) bool { return p[i].Val > p[j].Val }
func (p KVOrderDescs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// MapTopValN 找寻map[string]int中int最高的n个值
func MapTopValN(m map[string]int, n int) []string {
	out := make([]string, 0, n)
	kvs := make(KVOrderDescs, 0, len(m))
	for k, v := range m {
		kvs = append(kvs, KVOrderDesc{Key: k, Val: v})
	}
	sort.Sort(kvs)
	for i := 0; i < n && i < len(kvs); i++ {
		out = append(out, kvs[i].String())
	}
	return out
}

// MapTopValNInt 找寻map[int]int中int最高的n个值
func MapTopValNInt(m map[int]int, n int) []int {
	out := make([]int, 0, n)
	kvs := make(KVOrderDescs, 0, len(m))
	for k, v := range m {
		kvs = append(kvs, KVOrderDesc{Key: k, Val: v})
	}
	sort.Sort(kvs)
	for i := 0; i < n && i < len(kvs); i++ {
		out = append(out, kvs[i].Int())
	}
	return out
}

// Map2SliceDes Map2SliceDes
func Map2SliceDes(m map[string]int, fieldKey, fieldVal string) []map[string]interface{} {
	out := make([]map[string]interface{}, 0, len(m))
	kvs := make(KVOrderDescs, 0, len(m))
	for k, v := range m {
		kvs = append(kvs, KVOrderDesc{Key: k, Val: v})
	}
	sort.Sort(kvs)
	for i := 0; i < len(kvs); i++ {
		out = append(out, map[string]interface{}{
			fieldKey: kvs[i].Key,
			fieldVal: kvs[i].Val,
		})
	}
	return out
}

// MapKeyIncr get incr key slice
func MapKeyIncr(m map[string]interface{}) []string {
	return []string{}
}

// MapSSKeyIncr
func MapSSKeyIncr(m map[string]string) []string {
	ks := make([]string, 0)
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}

// MapSBKeyIncr
func MapSBKeyIncr(m map[string]bool) []string {
	ks := make([]string, 0)
	for k := range m {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	return ks
}

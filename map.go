package wtf

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

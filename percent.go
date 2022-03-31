package tf

// 百分比处理规则
// 因为凡有的，还要加给他，叫他有余；没有的，连他所有的也要夺过来。
// 如若有多个最大，则幸运留给第一个人。

func PercentInts(nums []int) []int {
	percent := make([]int, len(nums))
	sum := 0
	max := 0
	for _, num := range nums {
		sum += num
		if num > max {
			max = num
		}
	}
	for idx, num := range nums {
		percent[idx] = (num * 100) / sum
	}
	return percent
}

package wtf

import (
	"fmt"
	"time"
)

// TimeDesSechms 从秒转成XXhXXmXXs
func TimeDesSechms(i int) string {
	var (
		h, m, s int
	)
	s = i % 60
	m = ((i - s) / 60) % 60
	h = (i - m*60 - s) / 3600
	return fmt.Sprintf("%02dh%02dm%02d", h, m, s)
}

// TimeDesSecms 转换成 xxmxxs
func TimeDesSecms(i int) string {
	var (
		m, s int
	)
	s = i % 60
	m = i / 60
	return fmt.Sprintf("%02dm%02ds", m, s)
}

// TimeDesSecmsZh 转换成 xx分xx秒
func TimeDesSecmsZh(i int) string {
	var (
		m, s int
	)
	s = i % 60
	m = i / 60
	return fmt.Sprintf("%02d分%02d秒", m, s)
}

// TimeNow 返回现在时间2006-01-02 15:04:05
func TimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// TimeNowWithoutConnec 返回现在时间的20060102150405
func TimeNowWithoutConnec() string {
	return time.Now().Format("20060102150405")
}

// TimeDays 0 meaning today -1 is yesterday
func TimeDays(days ...int) string {
	if len(days) == 0 {
		days = []int{0}
	}
	return time.Now().AddDate(0, 0, days[0]).Format("2006-01-02")
}

// TimeMonths 返回这个月的时间2006-01
func TimeMonths(months ...int) string {
	if len(months) == 0 {
		months = []int{0}
	}
	return time.Now().AddDate(0, months[0], 0).Format("2006-01")
}

// TimeSplitDay 2017-01-10 15:54:32 return 2017-01-10
func TimeSplitDay(t string) string {
	if len(t) >= 10 {
		return t[:10]
	}
	return ""
}

// TimeSplitHMS 2017-01-10 15:54:32 return 15:54:32
func TimeSplitHMS(t string) string {
	if len(t) >= 19 {
		return t[11:]
	}
	return ""
}

// TimeSplitDayHMS return separately day & hms
func TimeSplitDayHMS(t string) (string, string) {
	return TimeSplitDay(t), TimeSplitHMS(t)
}

// TimePassSec 传入一个时间，返回距离现在过去了多少秒。
func TimePassSec(oldTime string) (int, string) {
	ot, _ := time.ParseInLocation("2006-01-02 15:04:05", oldTime, time.Local)
	now := time.Now()
	return int(now.Sub(ot).Seconds()), now.Format("2006-01-02 15:04:05")
}

// TimePassSecInt 只返回过去了多少秒。
func TimePassSecInt(oldTime string) int {
	ot, _ := time.ParseInLocation("2006-01-02 15:04:05", oldTime, time.Local)
	now := time.Now()
	return int(now.Sub(ot).Seconds())
}

// TimePassSecFrom 从指定的现在时间到旧时间经过了多少秒
func TimePassSecFrom(oldTime, nowTime string) int {
	ot, err := time.ParseInLocation("2006-01-02 15:04:05", oldTime, time.Local)
	if err != nil {
		return 0
	}
	nt, err := time.ParseInLocation("2006-01-02 15:04:05", nowTime, time.Local)
	if err != nil {
		return 0
	}
	return int(nt.Sub(ot).Seconds())
}

// MonthSpanToDay 将月-月 转换成相对应的天-天
func MonthSpanToDay(startMonth, endMonth string) (string, string) {
	if startMonth == "" {
		startMonth = TimeMonths()
	}
	if endMonth == "" {
		endMonth = startMonth
	}
	st, _ := time.ParseInLocation("2006-01", startMonth, time.Local)
	et, _ := time.ParseInLocation("2006-01", endMonth, time.Local)
	et = et.AddDate(0, 1, 0)
	et = et.Add(-1 * time.Second)
	return st.Format("2006-01-02"), et.Format("2006-01-02")
}

// TimeAddDays 给时间字符串增加天数
func TimeAddDays(t string, days int) string {
	ot, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	ot = ot.AddDate(0, 0, days)
	return ot.Format("2006-01-02 15:04:05")
}

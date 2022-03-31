package tf

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	TimeZHFormat string = "2006-01-02 15:04:05"
)

var (
	ZeroTime, _ = time.ParseInLocation("15:04:05", "00:00:00", time.Local)
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

// TimeDesSechmsZh 从秒转成XX小时XX分XX秒
func TimeDesSechmsZh(i int) string {
	var (
		h, m, s int
	)
	s = i % 60
	m = ((i - s) / 60) % 60
	h = (i - m*60 - s) / 3600
	return fmt.Sprintf("%02d小时%02d分%02d秒", h, m, s)
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

func Timestamp() int {
	return int(time.Now().Unix())
}

// TimeYMD 返回时间格式化为20060102的字符串
func TimeYMD() string {
	return time.Now().Format("20060102")
}

// TimeNowWithoutConnec 返回现在时间的20060102150405
func TimeNowWithoutConnec() string {
	return time.Now().Format("20060102150405")
}

// TimeNowUnixNano unixNano
func TimeNowUnixNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}

// TimeDayAddDays 在天字符串上加上天数
func TimeDayAddDays(day string, days int) string {
	t, _ := TimeParse(day)
	return t.AddDate(0, 0, days).Format("2006-01-02")
}

// TimeDayWeekday 那一天是星期几
func TimeDayWeekday(day string) int {
	t, _ := TimeParse(day)
	weekday := int(t.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	return weekday
}

// TimeDays 0 meaning today -1 is yesterday, format 2006-01-02
func TimeDays(days ...int) string {
	if len(days) == 0 {
		days = []int{0}
	}
	return time.Now().AddDate(0, 0, days[0]).Format("2006-01-02")
}

// TimeMonths 返回这个月的时间, format 2006-01
func TimeMonths(months ...int) string {
	if len(months) == 0 {
		months = []int{0}
	}
	return time.Now().AddDate(0, months[0], 0).Format("2006-01")
}

// TimeMonthsDays 返回这个月有多少天
func TimeMonthsDays(months ...int) int {
	if len(months) == 0 {
		months = []int{0}
	}
	t := TimeParseMust(time.Now().AddDate(0, months[0], 0).Format("2006-01"))
	return t.AddDate(0, 1, 0).Add(-1 * time.Minute).Day()
}

// TimeYears format 2006
func TimeYears(years ...int) string {
	if len(years) == 0 {
		years = []int{0}
	}
	return time.Now().AddDate(years[0], 0, 0).Format("2006")
}

// TimeMonthsFull format 当月第一天0秒
func TimeMonthsFull(months ...int) string {
	if len(months) == 0 {
		months = []int{0}
	}
	return time.Now().AddDate(0, months[0], 0).Format("2006-01") + "-01 00:00:00"
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
	ot, err := timeParse(oldTime)
	if err != nil {
		return 0
	}
	nt, err := timeParse(nowTime)
	if err != nil {
		return 0
	}
	return int(nt.Sub(ot).Seconds())
}

// TimePassDayFrom 从指定的现在时间到旧时间经过了多少天
func TimePassDayFrom(oldTime, nowTime string) int {
	ot, err := timeParse(oldTime)
	if err != nil {
		return 0
	}
	nt, err := timeParse(nowTime)
	if err != nil {
		return 0
	}
	return int(nt.Sub(ot).Hours() / 24)
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
	ot, _ := TimeParse(t)
	//ot, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	ot = ot.AddDate(0, 0, days)
	return ot.Format("2006-01-02 15:04:05")
}

// DurationsAvg []time.Duration 的平均值
func DurationsAvg(ds []time.Duration) time.Duration {
	sum := time.Duration(0)
	for _, d := range ds {
		sum += d
	}
	avg := sum / time.Duration(len(ds))
	return avg
}

// DurationsMaxMinAvg 返回 []time.Duration 最大值，最小值，平均值
func DurationsMaxMinAvg(ds []time.Duration) (max time.Duration, min time.Duration, avg time.Duration) {
	min = 1<<63 - 1
	sum := time.Duration(0)
	for _, d := range ds {
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
		sum += d
	}
	avg = sum / time.Duration(len(ds))
	return
}

// TimeParse 将字符串转换成时间
func TimeParse(v string) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	l := len(v)
	if l > len(layout) {
		l = len(layout)
		v = v[:len(layout)]
	}
	return time.ParseInLocation(layout[:len(v)], v, time.Local)
}

// Time 将字符串转成时间
func Time(v string) time.Time {
	t, _ := timeParse(v)
	return t
}

// TimeParseMust 字符串转时间
func TimeParseMust(v string) time.Time {
	t, _ := timeParse(v)
	return t
}

// TimeParseDay 将传入的时间转换成 2016-01-01，默认返回今天时间。
func TimeParseDay(s string) string {
	if len(s) < 10 {
		return TimeDays()
	}
	return s[:10]
}

const (
	timeFormat = "2006-01-02 15:04:05"

	stPadding = "2000-01-01 00:00:00"
	etPadding = "2999-12-31 23:59:59"
)

var (
	errPeriodParse = errors.New("period parse err")
)

// timeParse time parse from string
func timeParse(ts string) (time.Time, error) {
	format := timeFormat[:len(ts)]
	t, err := time.ParseInLocation(format, ts, time.Local)
	return t, err
}

// TimePeriodParse TimePeriodParse [)
func TimePeriodParse(st, et string) (string, string, error) {
	stl := len(st)
	etl := len(et)
	if stl == 0 {
		return "", "", errPeriodParse
	}
	sp := st + stPadding[stl:]
	var ep string
	spt, _ := timeParse(sp)
	if etl == 0 {
		bt := spt
		switch stl {
		case 4:
			ep = bt.AddDate(1, 0, 0).Format(timeFormat)
		case 7:
			ep = bt.AddDate(0, 1, 0).Format(timeFormat)
		case 10:
			ep = bt.AddDate(0, 0, 1).Format(timeFormat)
		case 13:
			ep = bt.Add(1 * time.Hour).Format(timeFormat)
		case 16:
			ep = bt.Add(1 * time.Minute).Format(timeFormat)
		case 19:
			ep = bt.Add(1 * time.Second).Format(timeFormat)
		default:
			return "", "", errPeriodParse
		}
	} else {
		if stl != etl {
			return "", "", errPeriodParse
		}
		ept, _ := timeParse(et)
		bt := ept
		switch stl {
		case 4:
			ep = bt.AddDate(1, 0, 0).Format(timeFormat)
		case 7:
			ep = bt.AddDate(0, 1, 0).Format(timeFormat)
		case 10:
			ep = bt.AddDate(0, 0, 1).Format(timeFormat)
		case 13:
			ep = bt.Add(1 * time.Hour).Format(timeFormat)
		case 16:
			ep = bt.Add(1 * time.Minute).Format(timeFormat)
		case 19:
			ep = bt.Add(1 * time.Second).Format(timeFormat)
		default:
			return "", "", errPeriodParse
		}
	}
	return sp, ep, nil
}

// IsWeekend is weekend
func IsWeekend(wd time.Weekday) bool {
	if wd == 0 || wd == 6 {
		return true
	}
	return false
}

// TimeDesAgo time des ago
func TimeDesAgo(t string) string {
	sec, _ := TimePassSec(t)
	if sec < 60 {
		return String(sec) + "秒前"
	} else if sec < 3600 {
		return String(sec/60) + "分钟前"
	} else if sec < 86400 {
		return String(sec/3600) + "小时前"
	}
	return String(sec/86400) + "天前"
}

// TimeUnixMillisecond Millisecond unix parse
func TimeUnixMillisecond(unix int64) time.Time {
	t := time.Unix(unix/1000, 0)
	t = t.Add(time.Duration(unix%1000) * time.Millisecond)
	return t
}

// TimeThisMon 这周周一
func TimeThisMon() string {
	t := time.Now()
	sub := int(t.Weekday()+6) % 7
	return t.AddDate(0, 0, -sub).Format("2006-01-02")
}

// TimeThisWeek 这个星期的所有日期
func TimeThisWeek() []string {
	t := time.Now()
	sub := int(t.Weekday()+6) % 7
	t = t.AddDate(0, 0, -sub)
	days := []string{}
	for i := 0; i < 7; i++ {
		days = append(days, t.AddDate(0, 0, i).Format("2006-01-02"))
	}
	return days
}

// TimeMonday 返回所在天的周一日期
// 1 <- 0
// 2 <- 1
// 7 <- 6
func TimeMonday(day string) string {
	t := Time(day)
	sub := int(t.Weekday()+6) % 7
	t = t.AddDate(0, 0, -sub)
	return t.Format("2006-01-02")
}

// TimeThisMonthDayLisToNow 本月日期合集（仅到本天为止）
func TimeThisMonthDayLisToNow() []string {
	month := TimeMonths()
	lis := []string{}
	start, err := timeParse(month)
	if err != nil {
		return lis
	}
	today := TimeDays()
	for ; start.Format("2006-01-02") <= today; start = start.AddDate(0, 0, 1) {
		lis = append(lis, start.Format("2006-01-02"))
	}
	return lis
}

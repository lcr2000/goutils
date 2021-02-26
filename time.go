package goutils

import (
	"github.com/spf13/cast"
	"time"
)

// TimeFormatter为时间格式化函数.
// i请传入time、int、int32、int64类型.
// format为需要格式化的日期格式 如：2006-01-02 15:04:05、2006.01.02等等.
func TimeFormatter(i interface{}, format string) string {
	var res time.Time
	switch i.(type) {
	case time.Time:
		res = cast.ToTime(i)
	case int, int32, int64:
		_t := cast.ToInt64(i)
		if _t == 0 {
			return ""
		}
		res = time.Unix(_t, 0)
	default:
		return ""
	}
	if res.IsZero() {
		return ""
	}
	return res.Format(format)
}

const defaultLayout = "2006-01-02"

// GetTodayStartTime获取本日起始时间.
// formatLayout为格式化格式.
func GetTodayStartTime(formatLayout string) time.Time {
	timeStr := time.Now().Format(formatLayout)
	t, _ := time.ParseInLocation(formatLayout, timeStr, time.Local)
	return t
}

// IsYesterday判断传入的时间t是否为昨天.
func IsYesterday(t int64) bool {
	t1 := GetTodayStartTime(defaultLayout)
	if t >= t1.AddDate(0, 0, -1).Unix() && t < t1.Unix() {
		return true
	}
	return false
}

// IsTomorrow判断传入的时间t是否为明天.
func IsTomorrow(t int64) bool {
	t1 := GetTodayStartTime(defaultLayout)
	if t >= t1.AddDate(0, 0, 1).Unix() && t < t1.AddDate(0, 0, 2).Unix() {
		return true
	}
	return false
}

// IsToday判断传入的时间t是否为今天.
func IsToday(t int64) bool {
	t1 := GetTodayStartTime(defaultLayout)
	if t >= t1.Unix() && t < t1.AddDate(0, 0, 1).Unix() {
		return true
	}
	return false
}
package times

import (
	"fmt"
	"king.com/king/rpc-ums/pkg/constants"
	"time"
)

const DATE_TIME = "2006-01-02 15:04:05"
const DATE_TIME_JOIN = "20060102150405999999"
const DATE = "2006-01-02"

func DateTimeFormat(dateTime time.Time, format string) string {
	if dateTime.IsZero() {
		return ""
	}
	return time.Time(dateTime).Format(format)
}
func DateDefaultFormat(date time.Time) string {
	if date.IsZero() {
		return ""
	}
	return DateTimeFormat(date, DATE_TIME)
}
func TimeToDateTime(t int64) string {
	return TimeToFormatStr(t, constants.DATE_TIME_FORMAT)
}

// 指定时间格式
func TimeToFormatStr(t int64, format string) string {
	return TimeToUnix(t).Format(format)
}

// 时间戳转time.Time
func TimeToUnix(t int64) time.Time {
	return time.Unix(t, 0)
}

// 当前时间
func CurrentDateTime() string {
	return TimeToDateTime(time.Now().Unix())
}

func StrToDefaultTime(dateStr string) (time.Time, error) {
	return StrToTime(dateStr, DATE_TIME)
}
func StrToTime(dateStr string, format string) (time.Time, error) {
	parsedTime, err := time.Parse(format, dateStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}, err
	}
	return parsedTime, nil
}

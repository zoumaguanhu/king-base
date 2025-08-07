package times

import (
	"fmt"
	"king.com/king/base/common/constants"
	"time"
)

// 时间格式
const DATE_TIME string = "2006-01-02 15:04:05"
const DATE_TIME_JOIN string = "20060102150405"
const DATE string = "2006-01-02"
const DATE_FORMAT = "2006-01-02"
const TIME_FORMAT = "15:04:05"

func DateTimeFormat(dateTime time.Time, format string) string {
	if dateTime.IsZero() {
		return constants.EMPTY_STRING
	}
	return dateTime.Format(format)
}
func DateDefaultFormat(date time.Time) string {
	if date.IsZero() {
		return constants.EMPTY_STRING
	}
	return DateTimeFormat(date, DATE_TIME)
}
func DateFormat(date time.Time) string {
	if date.IsZero() {
		return constants.EMPTY_STRING
	}
	return DateTimeFormat(date, DATE)
}
func TimeToDateTime(t int64) string {
	return TimeToFormatStr(t, DATE_TIME)
}
func TimeToDate(t int64) string {
	return TimeToFormatStr(t, DATE_FORMAT)
}

// TimeToFormatStr 指定时间格式
func TimeToFormatStr(t int64, format string) string {
	return TimeToUnix(t).Format(format)
}

// TimeToUnix 时间戳转time.Time
func TimeToUnix(t int64) time.Time {
	return time.Unix(t, 0)
}

// CurrentDateTime 当前时间
func CurrentDateTime() string {
	return TimeToDateTime(time.Now().Unix())
}
func CurrentDate() string {
	return TimeToDate(time.Now().Unix())
}

func StrToDefaultTime(dateStr string) (time.Time, error) {
	return StrToTime(dateStr, DATE_TIME)
}
func StrToTime(dateStr string, format string) (time.Time, error) {
	parsedTime, err := time.ParseInLocation(format, dateStr, time.Local)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}, err
	}
	return parsedTime, nil
}
func DateTimeToStartTime(t *time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 00, 00, 00, 0, t.Location())
}
func DateTimeToEndTime(t *time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}
func DateStrToEndTime(dateStr string) (time.Time, error) {
	toTime, err := StrToTime(dateStr, DATE)
	if err != nil {
		return time.Now(), err
	}
	return DateTimeToEndTime(&toTime), nil
}
func DateStrToStartTime(dateStr string) (time.Time, error) {
	toTime, err := StrToTime(dateStr, DATE)
	if err != nil {
		return time.Now(), err
	}
	return DateTimeToStartTime(&toTime), nil
}

func DateToFutureDate(d int64) *time.Time {
	n := time.Now()
	f := n.AddDate(0, 0, int(d))
	return &f
}
func ExpDurationTime(expTime string, t time.Duration) *time.Duration {
	endTime, err := StrToDefaultTime(expTime)
	if err != nil {
		return nil
	}
	d := endTime.Unix() - time.Now().Unix()
	v := time.Duration(d) * t
	return &v
}

// DateDiff 计算两个日期字符串之间的天数差
func DateDiff(dateStr1, dateStr2, layout string) (int, error) {
	// 解析日期字符串
	date1, err := time.Parse(layout, dateStr1)
	if err != nil {
		return 0, err
	}
	date2, err := time.Parse(layout, dateStr2)
	if err != nil {
		return 0, err
	}

	// 确保date2 >= date1，避免负数
	if date2.Before(date1) {
		date1, date2 = date2, date1
	}

	// 计算天数差（忽略时间部分，只计算日期差）
	day1 := date1.YearDay()
	day2 := date2.YearDay()
	year1 := date1.Year()
	year2 := date2.Year()

	// 如果年份不同，需要累加整年的天数
	if year1 != year2 {
		// 累加从year1到year2-1的整年天数
		for y := year1; y < year2; y++ {
			day2 += 365
			if isLeapYear(y) {
				day2++ // 闰年多一天
			}
		}
	}

	return day2 - day1, nil
}

// 判断是否为闰年
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
func PreviousDate() string {
	now := time.Now()
	yesterday := now.AddDate(0, 0, -1)

	return yesterday.Format(DATE) // Go的参考时间格式
}
func CurrentDayLastTime() int64 {
	now := time.Now()

	// 设置目标时间为当天的23点59分59秒
	targetTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())

	// 计算时间差
	diff := targetTime.Sub(now)

	return int64(diff.Seconds())
}

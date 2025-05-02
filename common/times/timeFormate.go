package times

import (
	"fmt"
	"king.com/king/rpc-ums/pkg/constants"
	"time"
)

const DATE_TIME string = "2006-01-02 15:04:05"
const DATE_TIME_JOIN string = "20060102150405"
const DATE string = "2006-01-02"

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
	return TimeToFormatStr(t, constants.DATE_TIME_FORMAT)
}
func TimeToDate(t int64) string {
	return TimeToFormatStr(t, constants.DATE_FORMAT)
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

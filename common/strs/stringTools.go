package strs

import (
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"king.com/king/base/common/arrs"
	"king.com/king/base/common/constants"
	"strconv"
	"strings"
)

// EmptyStr 空值
func EmptyStr() string {
	return constants.DEFAULT_STR
}

// IsEmpty 空判断
func IsEmpty(s *string) bool {
	return s == nil || *s == constants.DEFAULT_STR
}

// NotEmpty 非空判断
func NotEmpty(s *string) bool {
	return !IsEmpty(s)
}

// PStr2Val 指针判空
func PStr2Val(s *string) string {
	if s != nil {
		return *s
	}
	return constants.DEFAULT_STR
}

// Val2Ptr 字符串转指针
func Val2Ptr(s string) *string {
	if IsDefault(s) {
		return nil
	}
	return &s
}

// ArrayToString 数组转字符串
func ArrayToString(arr []string, n int) string {
	l := len(arr)
	if l == 0 {
		return constants.DEFAULT_STR
	}
	var sb strings.Builder
	sb.Grow(n)
	sb.WriteString(arr[0])
	if l > 1 {
		for _, s := range arr[1:] {
			sb.WriteString(constants.SEPARATOR_STR)
			sb.WriteString(s)
		}
	}

	return sb.String()
}
func ArrayTo100String(arr []string) string {
	return ArrayToString(arr, 100)
}
func ArrayTo24String(arr []string) string {
	return ArrayToString(arr, 24)
}

// StringToArray 字符串转数组
func StringToArray(s string) []string {
	if s == constants.DEFAULT_STR {
		return []string{}
	}
	return strings.Split(s, constants.SEPARATOR_STR)
}

// IsDefault 判断字符串是否默认值
func IsDefault(s string) bool {
	return s == constants.DEFAULT_STR
}
func NotDefault(s string) bool {
	return !IsDefault(s)
}
func ConvertUsd2Str(num int64) string {
	result := float64(num) / 100
	return fmt.Sprintf("%.2f", result)
}
func ConvertUsd2Float(num int64) float64 {
	result := float64(num) / 100
	temp := fmt.Sprintf("%.2f", result)
	ft, err := strconv.ParseFloat(temp, 64)
	if err != nil {
		return 0
	}
	return ft
}
func ConvertStr2Float(str string) float64 {
	result, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0
	}
	return result
}
func ConvertFloat2Str(str float64) string {
	return fmt.Sprintf("%.2f", str)
}
func StrToInt64Slice(s string, sp string) ([]int64, error) {
	parts := strings.Split(s, sp)
	result := make([]int64, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, num)
	}
	return result, nil
}
func StrsToInt64s(ss string) *[]int64 {
	arr := strings.Split(ss, constants.SEP_STR)
	r := &[]int64{}
	for _, id := range arr {
		d, _ := strconv.ParseInt(id, 10, 64)
		*r = append(*r, d)
	}
	return r
}
func StrsToInt64(s string) int64 {
	if IsDefault(s) {
		return constants.ZERO_INT64
	}
	d, _ := strconv.ParseInt(s, 10, 64)
	return d
}
func StrsToStr(ss []string) string {
	if arrs.IsDefault(ss) {
		return constants.DEFAULT_STR
	}
	return strings.Join(ss, constants.SEP_STR)
}
func StrsSplitToArr(s string, split string) *[]string {
	if IsDefault(s) {
		return &[]string{}
	}
	ss := strings.Split(s, split)
	return &ss
}
func StrsDefultSplitToArr(s string) *[]string {
	return StrsSplitToArr(s, constants.SEP_STR)
}
func GenPids(pIds string, id int64) string {
	return pIds + constants.SHORT_LINE + strconv.FormatInt(id, 10)
}
func ParsePids(pIds string) *[]int64 {
	return splitToInt64Array(pIds, constants.SHORT_LINE)
}
func splitToInt64Array(input string, splitFlg string) *[]int64 {
	// 按 "_" 分割字符串
	parts := strings.Split(input, splitFlg)
	numbers := &[]int64{}
	for _, part := range parts {
		// 将每个部分转换为 int64
		num, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			return numbers
		}
		*numbers = append(*numbers, num)
	}
	return numbers
}
func ParseSiteUrl(url string) *[]string {
	ss := &[]string{}
	arr := StringToArray(url)
	for _, h := range arr {
		cleaned := strings.TrimPrefix(h, constants.HTTPS_PREFIX)
		*ss = append(*ss, cleaned)
	}
	return ss
}
func ObjToStr(obj any) string {
	b, err := json.Marshal(obj)
	if err != nil {
		logx.Errorf("ObjToStr obj:%+v err:%v", obj, err)
		return constants.DEFAULT_STR
	}
	return string(b)
}
func StrToObj(s *string, c interface{}) {
	if err := json.Unmarshal([]byte(*s), c); err != nil {
		logx.Errorf("StrToObj obj:%v err:%v", *s, err)
	}

}
func IntToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}
func StrToFloat64(s string) float64 {
	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		logx.Errorf("StrToFloat64 s:%v err:%v", s, err)
	}
	return f64
}

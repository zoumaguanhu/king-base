package strs

import (
	"fmt"
	"king.com/king/base/common/constants"
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

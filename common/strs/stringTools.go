package strs

import (
	"bytes"
	"king.com/king/base/common/constants"
	"strings"
)

// IsEmpty 空判断
func IsEmpty(s *string) bool {
	return s == nil || *s == ""
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

// ArrayToString 数组转字符串
func ArrayToString(arr []string) string {
	if len(arr) == 0 {
		return constants.DEFAULT_STR
	}
	var buffer bytes.Buffer
	buffer.WriteString(arr[0])
	for _, s := range arr[1:] {
		buffer.WriteString(constants.SEPARATOR_STR)
		buffer.WriteString(s)
	}
	return buffer.String()
}

// StringToArray 字符串转数组
func StringToArray(s string) []string {
	if s == constants.DEFAULT_STR {
		return []string{}
	}
	return strings.Split(s, constants.SEPARATOR_STR)
}

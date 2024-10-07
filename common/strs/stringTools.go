package strs

import "king.com/king/base/common/constants"

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

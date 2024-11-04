package nums

import "king.com/king/base/common/constants"

func IsDefault(num int64) bool {
	return num == 0
}
func NotDefault(num int64) bool {
	return !IsDefault(num)
}
func PtrIsNil(num *int64) bool {
	return num == nil
}
func PtrNotNil(num *int64) bool {
	return !PtrIsNil(num)
}
func PStr2Val(num *int64) int64 {
	if num != nil {
		return *num
	}
	return constants.ZERO_INT64
}
func Val2Ptr(num int64) *int64 {
	if num == constants.ZERO_INT64 {
		return nil
	}
	return &num
}

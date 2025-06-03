package nums

import (
	"king.com/king/base/common/constants"
	"king.com/king/base/common/strs"
	"strconv"
	"strings"
)

func IsDefault(num int64) bool {
	return num == 0
}
func NotDefault(num int64) bool {
	return !IsDefault(num)
}
func PtrIsNil(num *int64) bool {
	return num == nil || *num == constants.ZERO_INT64
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
func Float64ToInt64(f float64) int64 {
	return int64(f)
}

// Difference 函数用于获取数组 A 中不包含数组 B 的元素集合
func Difference(A, B *[]int64) *[]int64 {
	// 创建一个 map 用于存储数组 B 中的元素
	setB := make(map[int64]bool)
	for _, num := range *B {
		setB[num] = true
	}

	var result []int64
	// 遍历数组 A，检查元素是否在 map 中
	for _, num := range *A {
		if !setB[num] {
			result = append(result, num)
		}
	}
	return &result
}
func Int64ToStr(num int64) string {
	return strconv.FormatInt(num, 10)
}
func FloatIsDefault(num float64) bool {
	return num == 0
}
func FloatToStr(f float64) string {
	return strconv.FormatFloat(f, 'f', 2, 64)
}
func SplitPrice(price string) (string, string, string) {
	num := strs.StrToFloat64(price)
	numStr := strconv.FormatFloat(num, 'f', 2, 64)
	parts := strings.Split(numStr, ".")
	// 确保小数部分存在
	decimalPart := constants.ZERO_ZERO
	if len(parts) > 1 {
		decimalPart = parts[1]
	}

	return parts[0], constants.DOT, decimalPart
}

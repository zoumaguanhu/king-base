package encryption

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	// 定义字符集
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+[]{}|;:'\",.<>/?"
)

// 生成指定长度的随机密码
func generatePassword(length int) string {
	// 确保字符集不为空
	allChars := fmt.Sprintf("%v%v%v", lowercaseLetters, uppercaseLetters, digits)
	charSetLength := len(allChars)
	if charSetLength == 0 || length <= 0 {
		return ""
	}

	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 存储生成的密码
	var password strings.Builder
	for i := 0; i < length; i++ {
		// 随机选择一个字符
		randomIndex := rand.Intn(charSetLength)
		password.WriteByte(allChars[randomIndex])
	}

	return password.String()
}
func GenPasswd() string {
	return generatePassword(8)
}

package encryption

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	"io"
	"math/rand"
	"regexp"
	"time"
)

/** 加密方式 **/
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"

var regx, _ = regexp.Compile("(?=.*[A-Za-z])(?=.*[0-9])")

func Md5ByString(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		panic(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

func Md5ByBytes(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}

// SaltBcrypt 加盐hash
func SaltBcrypt(passwordText string) string {
	password := []byte(passwordText)
	hashedPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	return string(hashedPassword)
}

// ComparePassword  密码校验
func ComparePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// 验证密码复杂度
func ValidPwdComplex(pwd string) bool {
	c := len(pwd)
	if c < 5 || c > 19 {
		return false
	}
	return regx.MatchString(pwd)
}

// 生成md5
func GenMd5(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

// 比较Md5
func CompareMd5(data string, md5str string) bool {
	return GenMd5(data) == md5str
}

// 生成指定长度字符串
func GeneratePassword(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

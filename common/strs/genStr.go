package strs

import (
	"fmt"
	"king.com/king/base/common/constants"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var gt = &RandomGenerator{}

type RandomGenerator struct{}

func (rg *RandomGenerator) generateRandomFourDigits() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	randomNumber := r.Intn(10000)
	return fmt.Sprintf("%04d", randomNumber)
}
func (rg *RandomGenerator) generateRandomAlphanumeric(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	result := make([]byte, length)
	usedIndexes := make(map[int]bool)
	for i := 0; i < length; i++ {
		index := r.Intn(len(constants.LETTER_BYTES))
		for usedIndexes[index] {
			index = r.Intn(len(constants.LETTER_BYTES))
		}
		usedIndexes[index] = true
		result[i] = constants.LETTER_BYTES[index]
	}
	return string(result)
}
func (rg *RandomGenerator) fillNumFourStr(num int64) string {
	str := strconv.FormatInt(num, 10)
	var lastFourDigits string
	if len(str) < 4 {
		lastFourDigits = str
	} else {
		lastFourDigits = str[len(str)-4:]
	}

	if len(lastFourDigits) < 4 {
		padding := constants.DEFAULT_STR
		for i := 0; i < 4-len(lastFourDigits); i++ {
			padding += constants.ZERO_STR
		}
		lastFourDigits = padding + lastFourDigits
	}
	return lastFourDigits
}
func GenStr(l int) string {
	return gt.generateRandomAlphanumeric(l)
}
func GenSnStr() string {
	return gt.generateRandomAlphanumeric(32)
}
func GenAcCodeStr() string {
	s := gt.generateRandomAlphanumeric(12)
	return strings.ToUpper(s)
}
func GenOrder(num int64) string {
	now := time.Now()

	return fmt.Sprintf("8%04d%02d%02d%04d%02d%02d%02d", now.Year(), now.Month(), now.Day(), num%constants.TEN_THOUSAND_64, now.Hour(), now.Minute(), now.Second())
}
func GenRefundOrder(num int64) string {
	now := time.Now()
	return fmt.Sprintf("7%04d%02d%02d%04d%02d%02d%02d", now.Year(), now.Month(), now.Day(), num%constants.TEN_THOUSAND_64, now.Hour(), now.Minute(), now.Second())

}
func GenPaySerial(num int64) string {
	now := time.Now()
	return fmt.Sprintf("9%04d%02d%02d%04d%02d%02d%02d", now.Year(), now.Month(), now.Day(), num%constants.TEN_THOUSAND_64, now.Hour(), now.Minute(), now.Second())

}
func GenDefaultOrder() string {
	return GenOrder(4)
}

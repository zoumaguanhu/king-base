package strs

import (
	"fmt"
	"king.com/king/base/common/constants"
	"math/rand"
	"time"
)

type RandomGenerator struct{}

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
func GenStr(l int) string {
	generator := &RandomGenerator{}
	return generator.generateRandomAlphanumeric(l)
}
func GenMno() string {
	now := time.Now()
	milliseconds := now.Nanosecond() / 1e6
	return fmt.Sprintf("8%04d%02d%02d%02d%02d%02d%03d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), milliseconds)
}

package encryption

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"king.com/king/base/common/constants"
)

// 生成指纹哈希
func GenerateFingerprint(ip, userAgent string) string {
	// 拼接IP和User-Agent
	data := fmt.Sprintf("%v%v%v", ip, constants.FINGER_FILL, userAgent)

	// 计算SHA256哈希
	hash := sha256.Sum256([]byte(data))

	// 转换为十六进制字符串
	return hex.EncodeToString(hash[:])
}

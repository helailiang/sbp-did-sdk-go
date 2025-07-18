package utils

import (
	"crypto/rand"
	"fmt"
)

// GenerateVCTemplateID 生成VC模板ID (SDK-011)
// 调用该方法将随机生成一个 UUIDv4 格式的字符串，作为凭证模版ID
func GenerateVCTemplateID() string {
	return generateUUIDv4()
}

// generateUUIDv4 生成UUIDv4格式的字符串
func generateUUIDv4() string {
	// 生成16字节的随机数
	bytes := make([]byte, 16)
	rand.Read(bytes)

	// 设置版本位 (version 4)
	bytes[6] = (bytes[6] & 0x0f) | 0x40

	// 设置变体位
	bytes[8] = (bytes[8] & 0x3f) | 0x80

	// 格式化为UUID字符串
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:16])
}

// ValidateUUID 验证UUID格式
func ValidateUUID(id string) bool {
	if len(id) != 36 {
		return false
	}

	// 检查格式: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	if id[8] != '-' || id[13] != '-' || id[18] != '-' || id[23] != '-' {
		return false
	}

	// 检查字符是否为十六进制
	for i, char := range id {
		if i == 8 || i == 13 || i == 18 || i == 23 {
			continue
		}
		if !isHexChar(char) {
			return false
		}
	}

	return true
}

// isHexChar 检查字符是否为十六进制字符
func isHexChar(char rune) bool {
	return (char >= '0' && char <= '9') ||
		(char >= 'a' && char <= 'f') ||
		(char >= 'A' && char <= 'F')
}

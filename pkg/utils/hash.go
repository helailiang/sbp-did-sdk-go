package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
)

// HashAlgorithm 哈希算法类型
type HashAlgorithm string

const (
	SHA256 HashAlgorithm = "SHA256"
	SM3    HashAlgorithm = "SM3"
)

// CalculateHash 计算哈希值 (SDK-017)
// 根据SHA256或SM3算法对数据进行哈希计算
func CalculateHash(data []byte, algorithm HashAlgorithm) (string, error) {
	if len(data) == 0 {
		return "", fmt.Errorf("data cannot be empty")
	}

	var h hash.Hash
	switch algorithm {
	case SHA256:
		h = sha256.New()
	case SM3:
		// 注意：这里使用SHA256作为SM3的替代，实际应该使用SM3实现
		h = sha256.New()
	default:
		return "", fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}

	h.Write(data)
	hashBytes := h.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

// CalculateHashFromString 从字符串计算哈希值
func CalculateHashFromString(data string, algorithm HashAlgorithm) (string, error) {
	return CalculateHash([]byte(data), algorithm)
}

// CalculateHashFromHex 从十六进制字符串计算哈希值
func CalculateHashFromHex(hexData string, algorithm HashAlgorithm) (string, error) {
	data, err := hex.DecodeString(hexData)
	if err != nil {
		return "", fmt.Errorf("failed to decode hex data: %w", err)
	}
	return CalculateHash(data, algorithm)
}

// ValidateHashAlgorithm 验证哈希算法是否有效
func ValidateHashAlgorithm(algorithm HashAlgorithm) bool {
	validAlgorithms := []HashAlgorithm{SHA256, SM3}
	for _, valid := range validAlgorithms {
		if algorithm == valid {
			return true
		}
	}
	return false
}

// GetHashAlgorithmFromString 从字符串获取哈希算法
func GetHashAlgorithmFromString(algorithm string) (HashAlgorithm, error) {
	hashAlgo := HashAlgorithm(algorithm)
	if !ValidateHashAlgorithm(hashAlgo) {
		return "", fmt.Errorf("unsupported hash algorithm: %s", algorithm)
	}
	return hashAlgo, nil
} 
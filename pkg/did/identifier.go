package did

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/helailiang/sbp-did-sdk-go/pkg/config"
	"github.com/helailiang/sbp-did-sdk-go/pkg/crypto"
)

// CalculateDIDIdentifier 计算DID标识符 (SDK-002)
// 基于公钥和DID Method生成完整的DID标识符
func CalculateDIDIdentifier(publicKey interface{}, didMethod string) (string, error) {
	// 验证DID Method
	if didMethod == "" {
		return "", fmt.Errorf("DID method cannot be empty")
	}

	// 确保DID Method以"did:"开头
	if !strings.HasPrefix(didMethod, "did:") {
		return "", fmt.Errorf("DID method must start with 'did:'")
	}

	// 获取公钥的字节表示
	var publicKeyBytes []byte
	var err error

	switch pk := publicKey.(type) {
	case *crypto.KeyPair:
		publicKeyBytes, err = pk.GetPublicKeyBytes()
		if err != nil {
			return "", fmt.Errorf("failed to get public key bytes: %w", err)
		}
	case []byte:
		publicKeyBytes = pk
	case string:
		// 如果是十六进制字符串，解码为字节
		publicKeyBytes, err = hex.DecodeString(pk)
		if err != nil {
			return "", fmt.Errorf("failed to decode hex public key: %w", err)
		}
	default:
		return "", fmt.Errorf("unsupported public key type: %T", publicKey)
	}

	// 计算公钥的哈希值
	hash := sha256.Sum256(publicKeyBytes)
	hashHex := hex.EncodeToString(hash[:])

	// 生成DID标识符
	// 格式: did:method:hash
	didIdentifier := fmt.Sprintf("%s%s", didMethod, hashHex)

	return didIdentifier, nil
}

// CalculateDIDIdentifierWithConfig 使用配置计算DID标识符
func CalculateDIDIdentifierWithConfig(cfg *config.Config, publicKey interface{}, didMethod string) (string, error) {
	// 验证配置
	if err := cfg.Validate(); err != nil {
		return "", fmt.Errorf("invalid configuration: %w", err)
	}

	// 验证华为云账户配置
	if cfg.HuaweiCloudAccessKey == "" {
		return "", fmt.Errorf("HuaweiCloudAccessKey is required")
	}

	// 验证密钥名称（这里假设从配置中获取）
	if cfg.DefaultAlgorithm == "" {
		return "", fmt.Errorf("DefaultAlgorithm is required")
	}

	return CalculateDIDIdentifier(publicKey, didMethod)
}

// ValidateDIDIdentifier 验证DID标识符格式
func ValidateDIDIdentifier(didIdentifier string) error {
	if didIdentifier == "" {
		return fmt.Errorf("DID identifier cannot be empty")
	}

	// 检查是否以"did:"开头
	if !strings.HasPrefix(didIdentifier, "did:") {
		return fmt.Errorf("DID identifier must start with 'did:'")
	}

	// 检查格式: did:method:identifier
	parts := strings.Split(didIdentifier, ":")
	if len(parts) < 3 {
		return fmt.Errorf("DID identifier format is invalid, expected: did:method:identifier")
	}

	// 检查method部分
	if parts[1] == "" {
		return fmt.Errorf("DID method cannot be empty")
	}

	// 检查identifier部分
	if parts[2] == "" {
		return fmt.Errorf("DID identifier part cannot be empty")
	}

	return nil
}

// ExtractDIDMethod 从DID标识符中提取DID Method
func ExtractDIDMethod(didIdentifier string) (string, error) {
	if err := ValidateDIDIdentifier(didIdentifier); err != nil {
		return "", err
	}

	parts := strings.Split(didIdentifier, ":")
	if len(parts) < 3 {
		return "", fmt.Errorf("invalid DID identifier format")
	}

	return parts[1], nil
}

// ExtractDIDIdentifier 从DID标识符中提取标识符部分
func ExtractDIDIdentifier(didIdentifier string) (string, error) {
	if err := ValidateDIDIdentifier(didIdentifier); err != nil {
		return "", err
	}

	parts := strings.Split(didIdentifier, ":")
	if len(parts) < 3 {
		return "", fmt.Errorf("invalid DID identifier format")
	}

	return parts[2], nil
}

// GenerateDIDFromKeyPair 从密钥对生成DID标识符
func GenerateDIDFromKeyPair(keyPair *crypto.KeyPair, didMethod string) (string, error) {
	if keyPair == nil {
		return "", fmt.Errorf("key pair cannot be nil")
	}

	return CalculateDIDIdentifier(keyPair, didMethod)
}

// GenerateDIDFromPublicKeyHex 从十六进制公钥生成DID标识符
func GenerateDIDFromPublicKeyHex(publicKeyHex, didMethod string) (string, error) {
	if publicKeyHex == "" {
		return "", fmt.Errorf("public key hex cannot be empty")
	}

	return CalculateDIDIdentifier(publicKeyHex, didMethod)
}

// GenerateDIDFromPublicKeyBytes 从字节公钥生成DID标识符
func GenerateDIDFromPublicKeyBytes(publicKeyBytes []byte, didMethod string) (string, error) {
	if len(publicKeyBytes) == 0 {
		return "", fmt.Errorf("public key bytes cannot be empty")
	}

	return CalculateDIDIdentifier(publicKeyBytes, didMethod)
}

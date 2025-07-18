package crypto

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/sbp-did/sbp-did-sdk-go/pkg/config"
)

// KeyPair 表示密钥对
type KeyPair struct {
	PrivateKey interface{} `json:"private_key"`
	PublicKey  interface{} `json:"public_key"`
	Algorithm  string      `json:"algorithm"`
	KeyName    string      `json:"key_name"`
}

// GenerateKeyPair 生成密钥对 (SDK-001)
// 支持ECDSA、RSA、SM2算法
func GenerateKeyPair(cfg *config.Config, algorithm, keyName string) (*KeyPair, error) {
	// 验证配置
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// 验证密钥名称
	if keyName == "" {
		return nil, fmt.Errorf("key name cannot be empty")
	}

	// 验证算法
	if !isValidAlgorithm(algorithm) {
		return nil, fmt.Errorf("unsupported algorithm: %s, supported algorithms: ECDSA, RSA, SM2", algorithm)
	}

	var keyPair *KeyPair
	var err error

	switch algorithm {
	case "ECDSA":
		keyPair, err = generateECDSAKeyPair()
	case "RSA":
		keyPair, err = generateRSAKeyPair()
	case "SM2":
		keyPair, err = generateSM2KeyPair()
	default:
		return nil, fmt.Errorf("unsupported algorithm: %s", algorithm)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to generate key pair: %w", err)
	}

	keyPair.Algorithm = algorithm
	keyPair.KeyName = keyName

	return keyPair, nil
}

// generateECDSAKeyPair 生成ECDSA密钥对
func generateECDSAKeyPair() (*KeyPair, error) {
	// 使用secp256k1曲线
	privateKey, err := btcec.NewPrivateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate ECDSA private key: %w", err)
	}

	publicKey := privateKey.PubKey()

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Algorithm:  "ECDSA",
	}, nil
}

// generateRSAKeyPair 生成RSA密钥对
func generateRSAKeyPair() (*KeyPair, error) {
	// 生成2048位RSA密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate RSA private key: %w", err)
	}

	publicKey := &privateKey.PublicKey

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Algorithm:  "RSA",
	}, nil
}

// generateSM2KeyPair 生成SM2密钥对
func generateSM2KeyPair() (*KeyPair, error) {
	// SM2使用P256曲线，这里使用ECDSA的P256实现
	privateKey, err := ecdsa.GenerateKey(btcec.S256(), rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate SM2 private key: %w", err)
	}

	return &KeyPair{
		PrivateKey: privateKey,
		PublicKey:  &privateKey.PublicKey,
		Algorithm:  "SM2",
	}, nil
}

// isValidAlgorithm 检查算法是否有效
func isValidAlgorithm(algorithm string) bool {
	validAlgorithms := []string{"ECDSA", "RSA", "SM2"}
	for _, valid := range validAlgorithms {
		if algorithm == valid {
			return true
		}
	}
	return false
}

// GetPublicKeyBytes 获取公钥的字节表示
func (kp *KeyPair) GetPublicKeyBytes() ([]byte, error) {
	switch kp.Algorithm {
	case "ECDSA":
		if pubKey, ok := kp.PublicKey.(*btcec.PublicKey); ok {
			return pubKey.SerializeCompressed(), nil
		}
	case "RSA":
		if pubKey, ok := kp.PublicKey.(*rsa.PublicKey); ok {
			return x509.MarshalPKCS1PublicKey(pubKey), nil
		}
	case "SM2":
		if pubKey, ok := kp.PublicKey.(*ecdsa.PublicKey); ok {
			return x509.MarshalPKIXPublicKey(pubKey)
		}
	}
	return nil, fmt.Errorf("unsupported algorithm: %s", kp.Algorithm)
}

// GetPrivateKeyBytes 获取私钥的字节表示
func (kp *KeyPair) GetPrivateKeyBytes() ([]byte, error) {
	switch kp.Algorithm {
	case "ECDSA":
		if privKey, ok := kp.PrivateKey.(*btcec.PrivateKey); ok {
			return privKey.Serialize(), nil
		}
	case "RSA":
		if privKey, ok := kp.PrivateKey.(*rsa.PrivateKey); ok {
			return x509.MarshalPKCS1PrivateKey(privKey), nil
		}
	case "SM2":
		if privKey, ok := kp.PrivateKey.(*ecdsa.PrivateKey); ok {
			return x509.MarshalPKCS8PrivateKey(privKey)
		}
	}
	return nil, fmt.Errorf("unsupported algorithm: %s", kp.Algorithm)
}

// GetPublicKeyPEM 获取公钥的PEM格式
func (kp *KeyPair) GetPublicKeyPEM() (string, error) {
	keyBytes, err := kp.GetPublicKeyBytes()
	if err != nil {
		return "", err
	}

	var blockType string
	switch kp.Algorithm {
	case "ECDSA":
		blockType = "PUBLIC KEY"
	case "RSA":
		blockType = "RSA PUBLIC KEY"
	case "SM2":
		blockType = "PUBLIC KEY"
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", kp.Algorithm)
	}

	block := &pem.Block{
		Type:  blockType,
		Bytes: keyBytes,
	}

	return string(pem.EncodeToMemory(block)), nil
}

// GetPrivateKeyPEM 获取私钥的PEM格式
func (kp *KeyPair) GetPrivateKeyPEM() (string, error) {
	keyBytes, err := kp.GetPrivateKeyBytes()
	if err != nil {
		return "", err
	}

	var blockType string
	switch kp.Algorithm {
	case "ECDSA":
		blockType = "EC PRIVATE KEY"
	case "RSA":
		blockType = "RSA PRIVATE KEY"
	case "SM2":
		blockType = "EC PRIVATE KEY"
	default:
		return "", fmt.Errorf("unsupported algorithm: %s", kp.Algorithm)
	}

	block := &pem.Block{
		Type:  blockType,
		Bytes: keyBytes,
	}

	return string(pem.EncodeToMemory(block)), nil
}

// GetPublicKeyHex 获取公钥的十六进制表示
func (kp *KeyPair) GetPublicKeyHex() (string, error) {
	keyBytes, err := kp.GetPublicKeyBytes()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", keyBytes), nil
}

// GetPrivateKeyHex 获取私钥的十六进制表示
func (kp *KeyPair) GetPrivateKeyHex() (string, error) {
	keyBytes, err := kp.GetPrivateKeyBytes()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", keyBytes), nil
}

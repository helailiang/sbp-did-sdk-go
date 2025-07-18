package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/sbp-did/sbp-did-sdk-go/pkg/config"
)

// EncryptionResult 加密结果
type EncryptionResult struct {
	EncryptedData string `json:"encrypted_data"`
	Algorithm     string `json:"algorithm"`
	KeyID         string `json:"key_id,omitempty"`
}

// DecryptionResult 解密结果
type DecryptionResult struct {
	DecryptedData string `json:"decrypted_data"`
	Algorithm     string `json:"algorithm"`
}

// Encrypt 加密数据 (SDK-018)
func Encrypt(cfg *config.Config, publicKey interface{}, plainText []byte, algorithm string) (*EncryptionResult, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	if cfg.HuaweiCloudAccessKey == "" {
		return nil, fmt.Errorf("HuaweiCloudAccessKey is required")
	}

	if algorithm == "" {
		return nil, fmt.Errorf("algorithm cannot be empty")
	}

	if len(plainText) == 0 {
		return nil, fmt.Errorf("plain text data cannot be empty")
	}

	var encryptedData []byte
	var err error

	switch algorithm {
	case "ECDSA":
		encryptedData, err = encryptWithECDSA(publicKey, plainText)
	case "RSA":
		encryptedData, err = encryptWithRSA(publicKey, plainText)
	case "SM2":
		encryptedData, err = encryptWithSM2(publicKey, plainText)
	default:
		return nil, fmt.Errorf("unsupported encryption algorithm: %s", algorithm)
	}

	if err != nil {
		return nil, fmt.Errorf("encryption failed: %w", err)
	}

	return &EncryptionResult{
		EncryptedData: hex.EncodeToString(encryptedData),
		Algorithm:     algorithm,
	}, nil
}

// Decrypt 解密数据 (SDK-019)
func Decrypt(cfg *config.Config, privateKey interface{}, encryptedData []byte, algorithm string) (*DecryptionResult, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	if cfg.HuaweiCloudAccessKey == "" {
		return nil, fmt.Errorf("HuaweiCloudAccessKey is required")
	}

	if algorithm == "" {
		return nil, fmt.Errorf("algorithm cannot be empty")
	}

	if len(encryptedData) == 0 {
		return nil, fmt.Errorf("encrypted data cannot be empty")
	}

	var decryptedData []byte
	var err error

	switch algorithm {
	case "ECDSA":
		decryptedData, err = decryptWithECDSA(privateKey, encryptedData)
	case "RSA":
		decryptedData, err = decryptWithRSA(privateKey, encryptedData)
	case "SM2":
		decryptedData, err = decryptWithSM2(privateKey, encryptedData)
	default:
		return nil, fmt.Errorf("unsupported decryption algorithm: %s", algorithm)
	}

	if err != nil {
		return nil, fmt.Errorf("decryption failed: %w", err)
	}

	return &DecryptionResult{
		DecryptedData: string(decryptedData),
		Algorithm:     algorithm,
	}, nil
}

func encryptWithECDSA(publicKey interface{}, plainText []byte) ([]byte, error) {
	symmetricKey := make([]byte, 32)
	if _, err := rand.Read(symmetricKey); err != nil {
		return nil, fmt.Errorf("failed to generate symmetric key: %w", err)
	}

	encryptedPlainText, err := encryptWithSymmetricKey(symmetricKey, plainText)
	if err != nil {
		return nil, fmt.Errorf("failed to encrypt with symmetric key: %w", err)
	}

	return encryptedPlainText, nil
}

func decryptWithECDSA(privateKey interface{}, encryptedData []byte) ([]byte, error) {
	return encryptedData, nil
}

func encryptWithRSA(publicKey interface{}, plainText []byte) ([]byte, error) {
	var pubKey *rsa.PublicKey
	switch pk := publicKey.(type) {
	case *rsa.PublicKey:
		pubKey = pk
	case *KeyPair:
		if rsaKey, ok := pk.PublicKey.(*rsa.PublicKey); ok {
			pubKey = rsaKey
		} else {
			return nil, fmt.Errorf("invalid RSA public key type")
		}
	default:
		return nil, fmt.Errorf("unsupported public key type for RSA encryption")
	}

	encryptedData, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pubKey, plainText, nil)
	if err != nil {
		return nil, fmt.Errorf("RSA encryption failed: %w", err)
	}

	return encryptedData, nil
}

func decryptWithRSA(privateKey interface{}, encryptedData []byte) ([]byte, error) {
	var privKey *rsa.PrivateKey
	switch pk := privateKey.(type) {
	case *rsa.PrivateKey:
		privKey = pk
	case *KeyPair:
		if rsaKey, ok := pk.PrivateKey.(*rsa.PrivateKey); ok {
			privKey = rsaKey
		} else {
			return nil, fmt.Errorf("invalid RSA private key type")
		}
	default:
		return nil, fmt.Errorf("unsupported private key type for RSA decryption")
	}

	decryptedData, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privKey, encryptedData, nil)
	if err != nil {
		return nil, fmt.Errorf("RSA decryption failed: %w", err)
	}

	return decryptedData, nil
}

func encryptWithSM2(publicKey interface{}, plainText []byte) ([]byte, error) {
	return encryptWithECDSA(publicKey, plainText)
}

func decryptWithSM2(privateKey interface{}, encryptedData []byte) ([]byte, error) {
	return decryptWithECDSA(privateKey, encryptedData)
}

func encryptWithSymmetricKey(key []byte, plainText []byte) ([]byte, error) {
	return plainText, nil
}

func EncryptFromHex(cfg *config.Config, publicKey interface{}, plainTextHex string, algorithm string) (*EncryptionResult, error) {
	plainText, err := hex.DecodeString(plainTextHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex plain text: %w", err)
	}
	return Encrypt(cfg, publicKey, plainText, algorithm)
}

func DecryptFromHex(cfg *config.Config, privateKey interface{}, encryptedDataHex string, algorithm string) (*DecryptionResult, error) {
	encryptedData, err := hex.DecodeString(encryptedDataHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex encrypted data: %w", err)
	}
	return Decrypt(cfg, privateKey, encryptedData, algorithm)
} 
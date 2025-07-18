package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/sbp-did/sbp-did-sdk-go/pkg/config"
)

// SignatureResult 签名结果
type SignatureResult struct {
	Signature string `json:"signature"`
	Algorithm string `json:"algorithm"`
	KeyID     string `json:"key_id,omitempty"`
}

// VerificationResult 验证结果
type VerificationResult struct {
	Valid     bool   `json:"valid"`
	Algorithm string `json:"algorithm"`
	Message   string `json:"message,omitempty"`
}

// Sign 签名数据 (SDK-020)
func Sign(cfg *config.Config, privateKey interface{}, data []byte, algorithm string) (*SignatureResult, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	if cfg.HuaweiCloudAccessKey == "" {
		return nil, fmt.Errorf("HuaweiCloudAccessKey is required")
	}

	if algorithm == "" {
		return nil, fmt.Errorf("algorithm cannot be empty")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data to sign cannot be empty")
	}

	var signature []byte
	var err error

	switch algorithm {
	case "ECDSA":
		signature, err = signWithECDSA(privateKey, data)
	case "RSA":
		signature, err = signWithRSA(privateKey, data)
	case "SM2":
		signature, err = signWithSM2(privateKey, data)
	default:
		return nil, fmt.Errorf("unsupported signature algorithm: %s", algorithm)
	}

	if err != nil {
		return nil, fmt.Errorf("signature failed: %w", err)
	}

	return &SignatureResult{
		Signature: hex.EncodeToString(signature),
		Algorithm: algorithm,
	}, nil
}

// VerifySignature 验证签名 (SDK-021)
func VerifySignature(cfg *config.Config, publicKey interface{}, data []byte, signature []byte, algorithm string) (*VerificationResult, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	if cfg.HuaweiCloudAccessKey == "" {
		return nil, fmt.Errorf("HuaweiCloudAccessKey is required")
	}

	if algorithm == "" {
		return nil, fmt.Errorf("algorithm cannot be empty")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("data to verify cannot be empty")
	}

	if len(signature) == 0 {
		return nil, fmt.Errorf("signature cannot be empty")
	}

	var valid bool
	var err error

	switch algorithm {
	case "ECDSA":
		valid, err = verifyWithECDSA(publicKey, data, signature)
	case "RSA":
		valid, err = verifyWithRSA(publicKey, data, signature)
	case "SM2":
		valid, err = verifyWithSM2(publicKey, data, signature)
	default:
		return nil, fmt.Errorf("unsupported verification algorithm: %s", algorithm)
	}

	if err != nil {
		return nil, fmt.Errorf("verification failed: %w", err)
	}

	result := &VerificationResult{
		Valid:     valid,
		Algorithm: algorithm,
	}

	if !valid {
		result.Message = "signature verification failed"
	}

	return result, nil
}

func signWithECDSA(privateKey interface{}, data []byte) ([]byte, error) {
	var privKey *ecdsa.PrivateKey
	switch pk := privateKey.(type) {
	case *ecdsa.PrivateKey:
		privKey = pk
	case *KeyPair:
		if ecdsaKey, ok := pk.PrivateKey.(*ecdsa.PrivateKey); ok {
			privKey = ecdsaKey
		} else {
			return nil, fmt.Errorf("invalid ECDSA private key type")
		}
	default:
		return nil, fmt.Errorf("unsupported private key type for ECDSA signing")
	}

	hash := sha256.Sum256(data)
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		return nil, fmt.Errorf("ECDSA signing failed: %w", err)
	}

	// 将r和s组合成签名
	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

func verifyWithECDSA(publicKey interface{}, data []byte, signature []byte) (bool, error) {
	var pubKey *ecdsa.PublicKey
	switch pk := publicKey.(type) {
	case *ecdsa.PublicKey:
		pubKey = pk
	case *KeyPair:
		if ecdsaKey, ok := pk.PublicKey.(*ecdsa.PublicKey); ok {
			pubKey = ecdsaKey
		} else {
			return false, fmt.Errorf("invalid ECDSA public key type")
		}
	default:
		return false, fmt.Errorf("unsupported public key type for ECDSA verification")
	}

	hash := sha256.Sum256(data)

	// 从签名中提取r和s
	if len(signature) < 64 {
		return false, fmt.Errorf("invalid signature length")
	}

	r := new(big.Int).SetBytes(signature[:32])
	s := new(big.Int).SetBytes(signature[32:64])

	valid := ecdsa.Verify(pubKey, hash[:], r, s)
	return valid, nil
}

func signWithRSA(privateKey interface{}, data []byte) ([]byte, error) {
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
		return nil, fmt.Errorf("unsupported private key type for RSA signing")
	}

	hash := sha256.Sum256(data)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, hash[:])
	if err != nil {
		return nil, fmt.Errorf("RSA signing failed: %w", err)
	}

	return signature, nil
}

func verifyWithRSA(publicKey interface{}, data []byte, signature []byte) (bool, error) {
	var pubKey *rsa.PublicKey
	switch pk := publicKey.(type) {
	case *rsa.PublicKey:
		pubKey = pk
	case *KeyPair:
		if rsaKey, ok := pk.PublicKey.(*rsa.PublicKey); ok {
			pubKey = rsaKey
		} else {
			return false, fmt.Errorf("invalid RSA public key type")
		}
	default:
		return false, fmt.Errorf("unsupported public key type for RSA verification")
	}

	hash := sha256.Sum256(data)
	err := rsa.VerifyPKCS1v15(pubKey, crypto.SHA256, hash[:], signature)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func signWithSM2(privateKey interface{}, data []byte) ([]byte, error) {
	// SM2签名实现，这里简化处理
	// 实际应该使用SM2算法
	return signWithECDSA(privateKey, data)
}

func verifyWithSM2(publicKey interface{}, data []byte, signature []byte) (bool, error) {
	// SM2验证实现，这里简化处理
	// 实际应该使用SM2算法
	return verifyWithECDSA(publicKey, data, signature)
}

func SignFromHex(cfg *config.Config, privateKey interface{}, dataHex string, algorithm string) (*SignatureResult, error) {
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex data: %w", err)
	}
	return Sign(cfg, privateKey, data, algorithm)
}

func VerifySignatureFromHex(cfg *config.Config, publicKey interface{}, dataHex string, signatureHex string, algorithm string) (*VerificationResult, error) {
	data, err := hex.DecodeString(dataHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex data: %w", err)
	}

	signature, err := hex.DecodeString(signatureHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode hex signature: %w", err)
	}

	return VerifySignature(cfg, publicKey, data, signature, algorithm)
}

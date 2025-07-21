package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/google/uuid"
)

// localKeyEntry 用于存储密钥及其类型
type localKeyEntry struct {
	keyType KeyType
	privKey interface{}
}

// LocalKeyManager 实现 Aries/TrustBloc 风格的本地 KeyManager 和 Crypto
type LocalKeyManager struct {
	mu    sync.Mutex
	store map[string]*localKeyEntry // keyID -> localKeyEntry
}

func NewLocalKeyManager() *LocalKeyManager {
	return &LocalKeyManager{
		store: make(map[string]*localKeyEntry),
	}
}

// Create 创建新密钥，返回 keyID 和公钥
func (l *LocalKeyManager) Create(keyType KeyType, opts ...KeyOpts) (string, []byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	keyID := uuid.NewString()
	var priv interface{}
	var pubBytes []byte
	var err error
	switch keyType {
	case ECDSAP256:
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			return "", nil, err
		}
		pubBytes, err = x509.MarshalPKIXPublicKey(&priv.(*ecdsa.PrivateKey).PublicKey)
	case RSA2048:
		priv, err = rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return "", nil, err
		}
		pubBytes, err = x509.MarshalPKIXPublicKey(&priv.(*rsa.PrivateKey).PublicKey)
	// 可扩展 Ed25519, SM2 ...
	default:
		return "", nil, fmt.Errorf("unsupported key type: %s", keyType)
	}
	if err != nil {
		return "", nil, err
	}
	l.store[keyID] = &localKeyEntry{keyType: keyType, privKey: priv}
	return keyID, pubBytes, nil
}

// Get 获取公钥
func (l *LocalKeyManager) Get(keyID string) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	entry, ok := l.store[keyID]
	if !ok {
		return nil, errors.New("key not found")
	}
	switch k := entry.privKey.(type) {
	case *ecdsa.PrivateKey:
		return x509.MarshalPKIXPublicKey(&k.PublicKey)
	case *rsa.PrivateKey:
		return x509.MarshalPKIXPublicKey(&k.PublicKey)
	default:
		return nil, errors.New("unsupported key type")
	}
}

// ImportPrivateKey 导入私钥，返回 keyID
func (l *LocalKeyManager) ImportPrivateKey(privKey []byte, keyType KeyType, opts ...KeyOpts) (string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	keyID := uuid.NewString()
	var priv interface{}
	var err error
	switch keyType {
	case ECDSAP256:
		priv, err = x509.ParseECPrivateKey(privKey)
	case RSA2048:
		priv, err = x509.ParsePKCS1PrivateKey(privKey)
	default:
		return "", fmt.Errorf("unsupported key type: %s", keyType)
	}
	if err != nil {
		return "", err
	}
	l.store[keyID] = &localKeyEntry{keyType: keyType, privKey: priv}
	return keyID, nil
}

// ExportPrivateKey 导出私钥
func (l *LocalKeyManager) ExportPrivateKey(keyID string) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	entry, ok := l.store[keyID]
	if !ok {
		return nil, errors.New("key not found")
	}
	switch k := entry.privKey.(type) {
	case *ecdsa.PrivateKey:
		return x509.MarshalECPrivateKey(k)
	case *rsa.PrivateKey:
		return x509.MarshalPKCS1PrivateKey(k), nil
	default:
		return nil, errors.New("unsupported key type")
	}
}

// Delete 删除密钥
func (l *LocalKeyManager) Delete(keyID string) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if _, ok := l.store[keyID]; !ok {
		return errors.New("key not found")
	}
	delete(l.store, keyID)
	return nil
}

// List 列举所有 keyID
func (l *LocalKeyManager) List() ([]string, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	ids := make([]string, 0, len(l.store))
	for id := range l.store {
		ids = append(ids, id)
	}
	return ids, nil
}

// Sign 使用指定 keyID 签名
func (l *LocalKeyManager) Sign(keyID string, data []byte) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	entry, ok := l.store[keyID]
	if !ok {
		return nil, errors.New("key not found")
	}
	hash := data
	switch k := entry.privKey.(type) {
	case *ecdsa.PrivateKey:
		r, s, err := ecdsa.Sign(rand.Reader, k, hash)
		if err != nil {
			return nil, err
		}
		return asn1.Marshal(struct{ R, S *big.Int }{r, s})
	case *rsa.PrivateKey:
		return rsa.SignPKCS1v15(rand.Reader, k, 0, hash)
	default:
		return nil, errors.New("unsupported key type")
	}
}

// Verify 使用指定 keyID 验签
func (l *LocalKeyManager) Verify(keyID string, data, signature []byte) (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	entry, ok := l.store[keyID]
	if !ok {
		return false, errors.New("key not found")
	}
	hash := data
	switch k := entry.privKey.(type) {
	case *ecdsa.PrivateKey:
		pub := &k.PublicKey
		var rs struct{ R, S *big.Int }
		_, err := asn1.Unmarshal(signature, &rs)
		if err != nil {
			return false, err
		}
		return ecdsa.Verify(pub, hash, rs.R, rs.S), nil
	case *rsa.PrivateKey:
		pub := &k.PublicKey
		err := rsa.VerifyPKCS1v15(pub, 0, hash, signature)
		return err == nil, nil
	default:
		return false, errors.New("unsupported key type")
	}
}

// Encrypt/Decrypt 可按需实现
func (l *LocalKeyManager) Encrypt(keyID string, plaintext []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
func (l *LocalKeyManager) Decrypt(keyID string, ciphertext []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

package crypto

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kms/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kms/v2/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"math/rand"
	"github.com/google/uuid"
)

// Aries/TrustBloc风格的华为KMS KeyManager和Crypto实现
//

// HuaweiKMSKeyManager 实现KeyManager和Crypto接口
//
type HuaweiKMSKeyManager struct {
	client    *kms.KmsClient
	projectId string
}

func NewHuaweiKMSKeyManager(endpoint, ak, sk, projectId string) (*HuaweiKMSKeyManager, error) {
	auth := basic.NewCredentialsBuilder().
		WithAk(ak).
		WithSk(sk).
		WithProjectId(projectId).
		Build()
	client := kms.NewKmsClient(
		kms.KmsClientBuilder().
			WithEndpoint(endpoint).
			WithCredential(auth).
			Build(),
	)
	return &HuaweiKMSKeyManager{client: client, projectId: projectId}, nil
}

// Create 创建新密钥，返回 keyID 和公钥
func (h *HuaweiKMSKeyManager) Create(keyType KeyType, opts ...KeyOpts) (string, []byte, error) {
	var keySpec string
	switch keyType {
	case RSA2048:
		keySpec = "RSA_2048"
	// 可扩展 ECDSA, SM2 ...
	default:
		return "", nil, fmt.Errorf("unsupported key type: %s", keyType)
	}
	keyAlias := "did-sdk-" + uuid.NewString()
	req := &model.CreateKeyRequest{
		Body: &model.CreateKeyRequestBody{
			KeyAlias: keyAlias,
			KeySpec:  keySpec,
			KeyUsage: "SIGN_VERIFY",
		},
	}
	resp, err := h.client.CreateKey(req)
	if err != nil {
		return "", nil, err
	}
	keyID := resp.KeyMetadata.KeyId
	// 获取公钥
	pub, err := h.Get(keyID)
	if err != nil {
		return "", nil, err
	}
	return keyID, pub, nil
}

// Get 获取公钥
func (h *HuaweiKMSKeyManager) Get(keyID string) ([]byte, error) {
	req := &model.ShowPublicKeyRequest{
		KeyId: keyID,
	}
	resp, err := h.client.ShowPublicKey(req)
	if err != nil {
		return nil, err
	}
	return []byte(resp.PublicKey), nil
}

// ImportPrivateKey KMS不支持导入私钥
func (h *HuaweiKMSKeyManager) ImportPrivateKey(privKey []byte, keyType KeyType, opts ...KeyOpts) (string, error) {
	return "", fmt.Errorf("Huawei KMS does not support importing private keys")
}

// ExportPrivateKey KMS不支持导出私钥
func (h *HuaweiKMSKeyManager) ExportPrivateKey(keyID string) ([]byte, error) {
	return nil, fmt.Errorf("Huawei KMS does not support exporting private keys")
}

// Delete 删除密钥
func (h *HuaweiKMSKeyManager) Delete(keyID string) error {
	req := &model.ScheduleKeyDeletionRequest{
		KeyId: keyID,
		Body: &model.ScheduleKeyDeletionRequestBody{PendingDays: 7}, // 7天后删除
	}
	_, err := h.client.ScheduleKeyDeletion(req)
	return err
}

// List 列举所有keyID（KMS不直接支持，需分页查询）
func (h *HuaweiKMSKeyManager) List() ([]string, error) {
	var ids []string
	marker := ""
	for {
		req := &model.ListKeysRequest{Marker: &marker, Limit: int32Ptr(100)}
		resp, err := h.client.ListKeys(req)
		if err != nil {
			return nil, err
		}
		for _, k := range *resp.Keys {
			ids = append(ids, k.KeyId)
		}
		if resp.NextMarker == nil || *resp.NextMarker == "" {
			break
		}
		marker = *resp.NextMarker
	}
	return ids, nil
}

func int32Ptr(i int32) *int32 { return &i }

// Sign 使用指定keyID签名
func (h *HuaweiKMSKeyManager) Sign(keyID string, data []byte) ([]byte, error) {
	req := &model.SignRequest{
		Body: &model.SignRequestBody{
			KeyId:           keyID,
			Message:         data,
			SigningAlgorithm: "RSASSA_PKCS1_V1_5_SHA_256",
			MessageType:     "DIGEST",
		},
	}
	resp, err := h.client.Sign(req)
	if err != nil {
		return nil, err
	}
	return []byte(resp.Signature), nil
}

// Verify 使用指定keyID验签
func (h *HuaweiKMSKeyManager) Verify(keyID string, data, signature []byte) (bool, error) {
	req := &model.VerifyRequest{
		Body: &model.VerifyRequestBody{
			KeyId:           keyID,
			Message:         data,
			Signature:       signature,
			SigningAlgorithm: "RSASSA_PKCS1_V1_5_SHA_256",
			MessageType:     "DIGEST",
		},
	}
	resp, err := h.client.Verify(req)
	if err != nil {
		return false, err
	}
	return *resp.Valid, nil
}

// Encrypt/Decrypt KMS暂不支持
func (h *HuaweiKMSKeyManager) Encrypt(keyID string, plaintext []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}
func (h *HuaweiKMSKeyManager) Decrypt(keyID string, ciphertext []byte) ([]byte, error) {
	return nil, fmt.Errorf("not implemented")
}

// RandString 生成随机字符串（用于KeyAlias）
func RandString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
} 
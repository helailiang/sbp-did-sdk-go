package did

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/helailiang/sbp-did-sdk-go/pkg/config"
	"github.com/helailiang/sbp-did-sdk-go/pkg/crypto"
)

// DIDDocument 表示DID文档结构（兼容TrustBloc/W3C）
type DIDDocument struct {
	Context            []string             `json:"@context"`
	ID                 string               `json:"id"`
	Controller         string               `json:"controller,omitempty"`
	VerificationMethod []VerificationMethod `json:"verificationMethod,omitempty"`
	Authentication     []string             `json:"authentication,omitempty"`
	AssertionMethod    []string             `json:"assertionMethod,omitempty"`
	KeyAgreement         []string               `json:"keyAgreement,omitempty"`
	CapabilityInvocation []string               `json:"capabilityInvocation,omitempty"`
	CapabilityDelegation []string               `json:"capabilityDelegation,omitempty"`
	Service              []Service              `json:"service,omitempty"`
	AlsoKnownAs          []string               `json:"alsoKnownAs,omitempty"`
	Created              string                 `json:"created,omitempty"`
	Updated              string                 `json:"updated,omitempty"`
	Deactivated          bool                   `json:"deactivated,omitempty"`
	CustomFields         map[string]interface{} `json:"-"`
}

// VerificationMethod 表示验证方法（兼容TrustBloc/W3C）
type VerificationMethod struct {
	ID              string      `json:"id"`
	Type            string      `json:"type"`
	Controller      string      `json:"controller"`
	PublicKeyBase58 string      `json:"publicKeyBase58,omitempty"`
	PublicKeyJwk    interface{} `json:"publicKeyJwk,omitempty"`
}

// PublicKeyJwk 表示JWK格式的公钥
type PublicKeyJwk struct {
	Kty string `json:"kty"`
	Crv string `json:"crv,omitempty"`
	X   string `json:"x,omitempty"`
	Y   string `json:"y,omitempty"`
	N   string `json:"n,omitempty"`
	E   string `json:"e,omitempty"`
	Kid string `json:"kid,omitempty"`
	Use string `json:"use,omitempty"`
}

// Service 表示服务端点
type Service struct {
	ID              string                 `json:"id"`
	Type            string                 `json:"type"`
	ServiceEndpoint string                 `json:"serviceEndpoint"`
	CustomFields    map[string]interface{} `json:"-"`
}

// AssembleDIDDocument 组装DID文档 (SDK-003)
// 需要公钥、算法、DID标识符和业务属性字段值
func AssembleDIDDocument(cfg *config.Config, publicKey interface{}, algorithm, didIdentifier string,
	businessAttributes map[string]interface{}) (*DIDDocument, error) {
	// 验证配置
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	// 验证华为云账户配置
	if cfg.HuaweiCloudAccessKey == "" {
		return nil, fmt.Errorf("HuaweiCloudAccessKey is required")
	}

	// 验证密钥名称
	if algorithm == "" {
		return nil, fmt.Errorf("algorithm cannot be empty")
	}

	// 验证DID标识符
	if didIdentifier == "" {
		return nil, fmt.Errorf("DID identifier cannot be empty")
	}

	// 验证DID标识符格式
	if err := ValidateDIDIdentifier(didIdentifier); err != nil {
		return nil, fmt.Errorf("invalid DID identifier: %w", err)
	}

	// 创建DID文档
	doc := &DIDDocument{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
			"https://w3id.org/security/suites/jws-2020/v1",
		},
		ID:           didIdentifier,
		Controller:   didIdentifier,
		Created:      time.Now().UTC().Format(time.RFC3339),
		Updated:      time.Now().UTC().Format(time.RFC3339),
		CustomFields: make(map[string]interface{}),
	}

	// 添加验证方法
	verificationMethod, err := createVerificationMethod(publicKey, algorithm, didIdentifier)
	if err != nil {
		return nil, fmt.Errorf("failed to create verification method: %w", err)
	}

	doc.VerificationMethod = []VerificationMethod{*verificationMethod}

	// 添加认证方法引用
	doc.Authentication = []string{verificationMethod.ID}

	// 添加断言方法引用
	doc.AssertionMethod = []string{verificationMethod.ID}

	// 添加业务属性字段
	if businessAttributes != nil {
		for key, value := range businessAttributes {
			doc.CustomFields[key] = value
		}
	}

	return doc, nil
}

// AssembleMultiKeyDIDDocument 组装包含多个密钥的DID文档
func AssembleMultiKeyDIDDocument(did string, keys []VerificationMethod, authKeys, assertionKeys []string) *DIDDocument {
	return &DIDDocument{
		Context:            []string{"https://www.w3.org/ns/did/v1"},
		ID:                 did,
		Controller:         did,
		VerificationMethod: keys,
		Authentication:     authKeys,
		AssertionMethod:    assertionKeys,
	}
}

// AddKey 向DID文档添加一个密钥，并指定用途（如authentication、assertionMethod等）
func (doc *DIDDocument) AddKey(newKey VerificationMethod, usages ...string) {
	doc.VerificationMethod = append(doc.VerificationMethod, newKey)
	for _, usage := range usages {
		switch usage {
		case "authentication":
			doc.Authentication = append(doc.Authentication, newKey.ID)
		case "assertionMethod":
			doc.AssertionMethod = append(doc.AssertionMethod, newKey.ID)
		case "keyAgreement":
			doc.KeyAgreement = append(doc.KeyAgreement, newKey.ID)
		case "capabilityInvocation":
			doc.CapabilityInvocation = append(doc.CapabilityInvocation, newKey.ID)
		case "capabilityDelegation":
			doc.CapabilityDelegation = append(doc.CapabilityDelegation, newKey.ID)
		}
	}
}

// RemoveKey 从DID文档中删除指定密钥（通过keyID），并同步移除所有用途中的引用
func (doc *DIDDocument) RemoveKey(keyID string) {
	// 移除VerificationMethod
	newVM := []VerificationMethod{}
	for _, vm := range doc.VerificationMethod {
		if vm.ID != keyID {
			newVM = append(newVM, vm)
		}
	}
	doc.VerificationMethod = newVM
	// 移除所有用途中的引用
	doc.Authentication = removeStringFromSlice(doc.Authentication, keyID)
	doc.AssertionMethod = removeStringFromSlice(doc.AssertionMethod, keyID)
	doc.KeyAgreement = removeStringFromSlice(doc.KeyAgreement, keyID)
	doc.CapabilityInvocation = removeStringFromSlice(doc.CapabilityInvocation, keyID)
	doc.CapabilityDelegation = removeStringFromSlice(doc.CapabilityDelegation, keyID)
}

// removeStringFromSlice 工具函数：从字符串切片中移除指定元素
func removeStringFromSlice(slice []string, target string) []string {
	result := []string{}
	for _, s := range slice {
		if s != target {
			result = append(result, s)
		}
	}
	return result
}

// createVerificationMethod 创建验证方法
func createVerificationMethod(publicKey interface{}, algorithm, didIdentifier string) (*VerificationMethod, error) {
	// 生成验证方法ID
	vmID := fmt.Sprintf("%s#keys-1", didIdentifier)

	// 获取公钥的不同格式
	var publicKeyHex string
	var publicKeyJwk *PublicKeyJwk
	var err error

	switch pk := publicKey.(type) {
	case *crypto.KeyPair:
		publicKeyHex, err = pk.GetPublicKeyHex()
		if err != nil {
			return nil, fmt.Errorf("failed to get public key hex: %w", err)
		}
		publicKeyJwk, err = createPublicKeyJwk(pk)
		if err != nil {
			return nil, fmt.Errorf("failed to create public key JWK: %w", err)
		}
	case string:
		publicKeyHex = pk
		// 对于字符串格式，尝试创建JWK
		publicKeyJwk, err = createPublicKeyJwkFromHex(pk, algorithm)
		if err != nil {
			// 如果失败，设置为nil
			publicKeyJwk = nil
		}
	default:
		return nil, fmt.Errorf("unsupported public key type: %T", publicKey)
	}

	// 确定验证方法类型
	var vmType string
	switch algorithm {
	case "ECDSA":
		vmType = "EcdsaSecp256k1VerificationKey2019"
	case "RSA":
		vmType = "RsaVerificationKey2018"
	case "SM2":
		vmType = "EcdsaSecp256k1VerificationKey2019" // SM2使用相同的类型
	default:
		vmType = "EcdsaSecp256k1VerificationKey2019"
	}

	vm := &VerificationMethod{
		ID:           vmID,
		Type:         vmType,
		Controller:   didIdentifier,
		PublicKeyHex: publicKeyHex,
		CustomFields: make(map[string]interface{}),
	}

	// 如果有JWK格式，添加它
	if publicKeyJwk != nil {
		vm.PublicKeyJwk = publicKeyJwk
	}

	return vm, nil
}

// createPublicKeyJwk 从密钥对创建JWK格式的公钥
func createPublicKeyJwk(keyPair *crypto.KeyPair) (*PublicKeyJwk, error) {
	switch keyPair.Algorithm {
	case "ECDSA":
		return createECDSAJwk(keyPair)
	case "RSA":
		return createRSAJwk(keyPair)
	case "SM2":
		return createSM2Jwk(keyPair)
	default:
		return nil, fmt.Errorf("unsupported algorithm for JWK: %s", keyPair.Algorithm)
	}
}

// createPublicKeyJwkFromHex 从十六进制字符串创建JWK
func createPublicKeyJwkFromHex(publicKeyHex, algorithm string) (*PublicKeyJwk, error) {
	// 这里简化实现，实际应该根据算法解析公钥
	switch algorithm {
	case "ECDSA":
		return &PublicKeyJwk{
			Kty: "EC",
			Crv: "secp256k1",
			X:   publicKeyHex[:64], // 简化的实现
			Y:   publicKeyHex[64:], // 简化的实现
		}, nil
	case "RSA":
		return &PublicKeyJwk{
			Kty: "RSA",
			N:   publicKeyHex, // 简化的实现
			E:   "AQAB",
		}, nil
	case "SM2":
		return &PublicKeyJwk{
			Kty: "EC",
			Crv: "secp256k1", // SM2使用相同的曲线表示
			X:   publicKeyHex[:64],
			Y:   publicKeyHex[64:],
		}, nil
	default:
		return nil, fmt.Errorf("unsupported algorithm: %s", algorithm)
	}
}

// createECDSAJwk 创建ECDSA的JWK
func createECDSAJwk(keyPair *crypto.KeyPair) (*PublicKeyJwk, error) {
	// 这里需要根据实际的ECDSA公钥实现
	// 简化实现
	return &PublicKeyJwk{
		Kty: "EC",
		Crv: "secp256k1",
		X:   "sample-x-coordinate",
		Y:   "sample-y-coordinate",
	}, nil
}

// createRSAJwk 创建RSA的JWK
func createRSAJwk(keyPair *crypto.KeyPair) (*PublicKeyJwk, error) {
	// 这里需要根据实际的RSA公钥实现
	// 简化实现
	return &PublicKeyJwk{
		Kty: "RSA",
		N:   "sample-modulus",
		E:   "AQAB",
	}, nil
}

// createSM2Jwk 创建SM2的JWK
func createSM2Jwk(keyPair *crypto.KeyPair) (*PublicKeyJwk, error) {
	// SM2使用与ECDSA相同的JWK格式
	return createECDSAJwk(keyPair)
}

// ToJSON 将DID文档转换为JSON
func (doc *DIDDocument) ToJSON() ([]byte, error) {
	// 创建包含自定义字段的map
	docMap := make(map[string]interface{})

	// 序列化基本字段
	basicJSON, err := json.Marshal(doc)
	if err != nil {
		return nil, err
	}

	// 解析基本JSON
	err = json.Unmarshal(basicJSON, &docMap)
	if err != nil {
		return nil, err
	}

	// 添加自定义字段
	for key, value := range doc.CustomFields {
		docMap[key] = value
	}

	// 重新序列化
	return json.MarshalIndent(docMap, "", "  ")
}

// FromJSON 从JSON创建DID文档
func FromJSON(data []byte) (*DIDDocument, error) {
	var doc DIDDocument
	err := json.Unmarshal(data, &doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

// ValidateDIDDocument 验证DID文档
func ValidateDIDDocument(doc *DIDDocument) error {
	if doc == nil {
		return fmt.Errorf("DID document cannot be nil")
	}

	if doc.ID == "" {
		return fmt.Errorf("DID document ID cannot be empty")
	}

	if err := ValidateDIDIdentifier(doc.ID); err != nil {
		return fmt.Errorf("invalid DID document ID: %w", err)
	}

	if len(doc.Context) == 0 {
		return fmt.Errorf("DID document context cannot be empty")
	}

	return nil
}

// NewVerificationMethodFromKeyManager 工具函数：通过KeyManager生成的密钥自动组装VerificationMethod
// 参数：didIdentifier、keyID、algorithm、keyManager
// 返回：VerificationMethod，error
func NewVerificationMethodFromKeyManager(didIdentifier, keyID, algorithm string, keyManager crypto.KeyManager) (*VerificationMethod, error) {
	pubKey, err := keyManager.GetPublicKey(keyID)
	if err != nil {
		return nil, err
	}
	// 组装JWK（这里只做简单示例，实际应根据算法和公钥类型填充x/y/n/e等字段）
	var jwk interface{}
	switch algorithm {
	case "ECDSA":
		jwk = map[string]interface{}{"kty": "EC", "crv": "secp256k1", "x": "x", "y": "y"} // TODO: 真实x/y
	case "RSA":
		jwk = map[string]interface{}{"kty": "RSA", "n": "n", "e": "AQAB"} // TODO: 真实n
	case "SM2":
		jwk = map[string]interface{}{"kty": "EC", "crv": "secp256k1", "x": "x", "y": "y"}
	default:
		jwk = nil
	}
	return &VerificationMethod{
		ID:           didIdentifier + "#" + keyID,
		Type:         map[string]string{"ECDSA": "EcdsaSecp256k1VerificationKey2019", "RSA": "RsaVerificationKey2018", "SM2": "EcdsaSecp256k1VerificationKey2019"}[algorithm],
		Controller:   didIdentifier,
		PublicKeyJwk: jwk,
	}, nil
}
 
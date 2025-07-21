package wallet

// Collection 表示钱包中的通用集合（如VC、DID、Key等的分组）
type Collection struct {
	ID          string            `json:"id"`          // 集合唯一标识
	Type        string            `json:"type"`        // 集合类型（如"credential"、"key"、"did"等）
	Name        string            `json:"name"`        // 集合名称
	Description string            `json:"description"` // 集合描述
	Owner       string            `json:"owner"`       // 所属用户DID
	Tags        map[string]string `json:"tags"`        // 可选标签
}

// Credential 表示可验证凭证（兼容W3C标准）
type Credential struct {
	ID                string                 `json:"id"`
	Context           []string               `json:"@context"`
	Type              []string               `json:"type"`
	Issuer            string                 `json:"issuer"`
	IssuanceDate      string                 `json:"issuanceDate"`
	ExpirationDate    string                 `json:"expirationDate,omitempty"`
	CredentialSubject map[string]interface{} `json:"credentialSubject"`
	CredentialSchema  interface{}            `json:"credentialSchema,omitempty"`
	Proof             *Proof                 `json:"proof,omitempty"`
}

// DIDResolutionResponse 表示DID解析响应（兼容W3C标准）
type DIDResolutionResponse struct {
	Context              interface{}   `json:"@context"`
	DidDocument          interface{}   `json:"didDocument"`
	DidDocumentMetadata  interface{}   `json:"didDocumentMetadata"`
	DidResolutionMetadata interface{}  `json:"didResolutionMetadata"`
}

// Key 表示钱包内的密钥元数据
// 仅存储元信息，私钥可由KeyManager管理
//
type Key struct {
	ID        string `json:"id"`        // keyID
	Type      string `json:"type"`      // 密钥类型（如ECDSA、RSA、Ed25519等）
	Controller string `json:"controller"` // 拥有者DID
	PublicKey  []byte `json:"publicKey"`  // 公钥
	Created    string `json:"created"`    // 创建时间
	Tags       map[string]string `json:"tags"` // 可选标签
}

// VerifiablePresentation 表示可验证表示（VP，兼容W3C标准）
type VerifiablePresentation struct {
	ID                string        `json:"id"`
	Context           []string      `json:"@context"`
	Type              []string      `json:"type"`
	Holder            string        `json:"holder"`
	VerifiableCredential []interface{} `json:"verifiableCredential"` // 可为Credential或其引用
	Proof             *Proof        `json:"proof,omitempty"`
}

// WalletBackup 表示钱包备份数据结构
// 可用于导出/恢复钱包全部数据
//
type WalletBackup struct {
	Users       []*WalletUser         `json:"users"`
	CreatedAt   string               `json:"createdAt"`
	BackupNotes string               `json:"backupNotes,omitempty"`
}

// Proof 兼容W3C标准的证明结构体
// https://www.w3.org/TR/vc-data-model/#proofs-0
type Proof struct {
	Type               string      `json:"type"`
	Created            string      `json:"created"`
	ProofPurpose       string      `json:"proofPurpose"`
	VerificationMethod string      `json:"verificationMethod"`
	Jws                string      `json:"jws,omitempty"`
	SignatureValue     string      `json:"signatureValue,omitempty"`
	Domain             string      `json:"domain,omitempty"`
	Challenge          string      `json:"challenge,omitempty"`
	CustomFields       interface{} `json:"customFields,omitempty"` // 兼容扩展
} 
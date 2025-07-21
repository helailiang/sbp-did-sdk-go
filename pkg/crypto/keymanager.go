package crypto

// KeyType 定义支持的密钥类型
// 参考 Aries/TrustBloc KMS 设计
// 可扩展 Ed25519, ECDSAP256, RSA2048, SM2 等
//
type KeyType string

const (
    ED25519   KeyType = "ED25519"
    ECDSAP256 KeyType = "ECDSAP256"
    RSA2048   KeyType = "RSA2048"
    SM2       KeyType = "SM2"
)

// KeyManager Aries/TrustBloc 风格接口
// 所有密钥用唯一 keyID 标识，支持多后端实现
//
type KeyManager interface {
    // Create 创建新密钥，返回 keyID 和公钥
    Create(keyType KeyType, opts ...KeyOpts) (keyID string, pubKey []byte, err error)
    // Get 获取公钥
    Get(keyID string) (pubKey []byte, err error)
    // ImportPrivateKey 导入私钥，返回 keyID
    ImportPrivateKey(privKey []byte, keyType KeyType, opts ...KeyOpts) (keyID string, err error)
    // ExportPrivateKey 导出私钥（可选，部分后端不支持）
    ExportPrivateKey(keyID string) ([]byte, error)
    // Delete 删除密钥
    Delete(keyID string) error
    // List 列举所有 keyID
    List() ([]string, error)
}

// KeyOpts 密钥用途、标签等扩展属性
// 可根据需要扩展
//
type KeyOpts interface {
    Purpose() string // e.g. "signing", "encryption"
}

// Crypto Aries/TrustBloc 风格接口
// 通过 keyID 进行签名、验签、加解密等操作
//
type Crypto interface {
    Sign(keyID string, data []byte) ([]byte, error)
    Verify(keyID string, data, signature []byte) (bool, error)
    Encrypt(keyID string, plaintext []byte) ([]byte, error)
    Decrypt(keyID string, ciphertext []byte) ([]byte, error)
} 
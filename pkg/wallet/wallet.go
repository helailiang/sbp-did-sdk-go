package wallet

import (
	"encoding/json"
	"errors"
	"sync"
	"time"

	"github.com/helailiang/sbp-did-sdk-go/pkg/crypto"
	"github.com/helailiang/sbp-did-sdk-go/pkg/did"
)

// WalletUser 表示一个钱包用户
// 每个用户可独立选择密钥后端（本地或KMS）
type WalletUser struct {
	DID        string
	DIDDoc     *did.DIDDocument // 新增：本地DID文档缓存
	KeyManager crypto.KeyManager
	Collections map[string]*Collection
	Credentials map[string]*Credential
	Keys        map[string]*Key
	Mutex       sync.RWMutex
}

// AddDIDKey 向本地DID文档添加密钥并指定用途
func (u *WalletUser) AddDIDKey(newKey did.VerificationMethod, usages ...string) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if u.DIDDoc == nil {
		return errors.New("DID document not initialized")
	}
	u.DIDDoc.AddKey(newKey, usages...)
	return nil
}

// RemoveDIDKey 从本地DID文档移除密钥
func (u *WalletUser) RemoveDIDKey(keyID string) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if u.DIDDoc == nil {
		return errors.New("DID document not initialized")
	}
	u.DIDDoc.RemoveKey(keyID)
	return nil
}

// SyncDIDDocument 将本地DID文档同步到链上（调用OpenAPI UpdateDID）
func (u *WalletUser) SyncDIDDocument(apiClient interface{}, projectNo, signature, txSignature string, index int) error {
	u.Mutex.RLock()
	defer u.Mutex.RUnlock()
	if u.DIDDoc == nil {
		return errors.New("DID document not initialized")
	}
	docJSON, err := u.DIDDoc.ToJSON()
	if err != nil {
		return err
	}
	// 断言apiClient为*api.Client
	client, ok := apiClient.(interface {
		UpdateDID(req interface{}) (interface{}, error)
	})
	if !ok {
		return errors.New("invalid api client type")
	}
	// 构造UpdateDIDRequest
	req := map[string]interface{}{
		"projectNo":   projectNo,
		"didDocument": string(docJSON),
		"index":       index,
		"signature":   signature,
		"txSignature": txSignature,
	}
	_, err = client.UpdateDID(req)
	return err
}

// Wallet 支持多用户和多后端密钥管理
//
type Wallet struct {
	users map[string]*WalletUser // DID -> WalletUser
	mutex sync.RWMutex
}

// NewWallet 创建新的钱包实例
func NewWallet() *Wallet {
	return &Wallet{
		users: make(map[string]*WalletUser),
	}
}

// AddUser 添加新用户，需指定DID和KeyManager
func (w *Wallet) AddUser(did string, keyManager crypto.KeyManager) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	if _, exists := w.users[did]; exists {
		return errors.New("user already exists")
	}
	w.users[did] = &WalletUser{
		DID:         did,
		KeyManager:  keyManager,
		Collections: make(map[string]*Collection),
		Credentials: make(map[string]*Credential),
		Keys:        make(map[string]*Key),
	}
	return nil
}

// GetUser 获取用户
func (w *Wallet) GetUser(did string) (*WalletUser, error) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	user, ok := w.users[did]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// ========== Collection 操作 ==========
func (u *WalletUser) AddCollection(col *Collection) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, exists := u.Collections[col.ID]; exists {
		return errors.New("collection already exists")
	}
	u.Collections[col.ID] = col
	return nil
}

func (u *WalletUser) GetCollection(id string) (*Collection, error) {
	u.Mutex.RLock()
	defer u.Mutex.RUnlock()
	col, ok := u.Collections[id]
	if !ok {
		return nil, errors.New("collection not found")
	}
	return col, nil
}

func (u *WalletUser) DeleteCollection(id string) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, ok := u.Collections[id]; !ok {
		return errors.New("collection not found")
	}
	delete(u.Collections, id)
	return nil
}

// ========== Credential 操作 ==========
func (u *WalletUser) AddCredential(vc *Credential) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, exists := u.Credentials[vc.ID]; exists {
		return errors.New("credential already exists")
	}
	u.Credentials[vc.ID] = vc
	return nil
}

func (u *WalletUser) GetCredential(id string) (*Credential, error) {
	u.Mutex.RLock()
	defer u.Mutex.RUnlock()
	vc, ok := u.Credentials[id]
	if !ok {
		return nil, errors.New("credential not found")
	}
	return vc, nil
}

func (u *WalletUser) DeleteCredential(id string) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, ok := u.Credentials[id]; !ok {
		return errors.New("credential not found")
	}
	delete(u.Credentials, id)
	return nil
}

// ========== Key 操作 ==========
func (u *WalletUser) AddKey(key *Key) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, exists := u.Keys[key.ID]; exists {
		return errors.New("key already exists")
	}
	u.Keys[key.ID] = key
	return nil
}

func (u *WalletUser) GetKey(id string) (*Key, error) {
	u.Mutex.RLock()
	defer u.Mutex.RUnlock()
	key, ok := u.Keys[id]
	if !ok {
		return nil, errors.New("key not found")
	}
	return key, nil
}

func (u *WalletUser) DeleteKey(id string) error {
	u.Mutex.Lock()
	defer u.Mutex.Unlock()
	if _, ok := u.Keys[id]; !ok {
		return errors.New("key not found")
	}
	delete(u.Keys, id)
	return nil
}

// ========== DIDResolutionResponse 操作 ==========
// 可根据业务需求扩展，通常解析结果不直接存储在钱包内

// ========== KeyManager 相关 ==========
// 可通过u.KeyManager进行密钥生命周期管理、签名、验签等操作

// ========== 辅助方法 ==========
// 可扩展导入/导出、备份恢复、标签搜索等高级功能 

// Backup 导出钱包所有用户数据
func (w *Wallet) Backup(notes string) (*WalletBackup, error) {
	w.mutex.RLock()
	defer w.mutex.RUnlock()
	users := make([]*WalletUser, 0, len(w.users))
	for _, u := range w.users {
		users = append(users, u)
	}
	return &WalletBackup{
		Users: users,
		CreatedAt: time.Now().Format(time.RFC3339),
		BackupNotes: notes,
	}, nil
}

// Restore 恢复钱包所有用户数据（覆盖原有数据）
func (w *Wallet) Restore(backup *WalletBackup) error {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.users = make(map[string]*WalletUser)
	for _, u := range backup.Users {
		w.users[u.DID] = u
	}
	return nil
}

// BackupToJSON 导出为JSON
func (w *Wallet) BackupToJSON(notes string) ([]byte, error) {
	backup, err := w.Backup(notes)
	if err != nil {
		return nil, err
	}
	return json.MarshalIndent(backup, "", "  ")
}

// RestoreFromJSON 从JSON恢复
func (w *Wallet) RestoreFromJSON(data []byte) error {
	var backup WalletBackup
	if err := json.Unmarshal(data, &backup); err != nil {
		return err
	}
	return w.Restore(&backup)
} 
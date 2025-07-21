package tests

import (
	"testing"

	"github.com/helailiang/sbp-did-sdk-go/pkg/config"
	"github.com/helailiang/sbp-did-sdk-go/pkg/crypto"
	"github.com/helailiang/sbp-did-sdk-go/pkg/did"
	"github.com/helailiang/sbp-did-sdk-go/pkg/wallet"
)

func TestWalletBasic(t *testing.T) {
	w := wallet.NewWallet()
	user1DID := "did:sbp:user1"
	user2DID := "did:sbp:user2"
	user1KM := crypto.NewLocalKeyManager()
	user2KM := crypto.NewLocalKeyManager()

	// 添加用户
	if err := w.AddUser(user1DID, user1KM); err != nil {
		t.Fatalf("AddUser user1 failed: %v", err)
	}
	if err := w.AddUser(user2DID, user2KM); err != nil {
		t.Fatalf("AddUser user2 failed: %v", err)
	}

	// 获取用户
	user1, err := w.GetUser(user1DID)
	if err != nil {
		t.Fatalf("GetUser user1 failed: %v", err)
	}
	user2, err := w.GetUser(user2DID)
	if err != nil {
		t.Fatalf("GetUser user2 failed: %v", err)
	}

	// 用户1添加VC
	vc := &wallet.Credential{ID: "vc-001", Issuer: user1DID}
	if err := user1.AddCredential(vc); err != nil {
		t.Fatalf("AddCredential failed: %v", err)
	}
	// 查询VC
	vcGot, err := user1.GetCredential("vc-001")
	if err != nil || vcGot.ID != "vc-001" {
		t.Fatalf("GetCredential failed: %v", err)
	}
	// 删除VC
	if err := user1.DeleteCredential("vc-001"); err != nil {
		t.Fatalf("DeleteCredential failed: %v", err)
	}
	if _, err := user1.GetCredential("vc-001"); err == nil {
		t.Fatalf("Credential should be deleted")
	}

	// 用户1添加Key
	keyID, pub, _ := user1KM.Create(crypto.ECDSAP256)
	key := &wallet.Key{ID: keyID, Type: "ECDSA", Controller: user1DID, PublicKey: pub}
	if err := user1.AddKey(key); err != nil {
		t.Fatalf("AddKey failed: %v", err)
	}
	keyGot, err := user1.GetKey(keyID)
	if err != nil || keyGot.ID != keyID {
		t.Fatalf("GetKey failed: %v", err)
	}
	if err := user1.DeleteKey(keyID); err != nil {
		t.Fatalf("DeleteKey failed: %v", err)
	}
	if _, err := user1.GetKey(keyID); err == nil {
		t.Fatalf("Key should be deleted")
	}

	// 用户1添加Collection
	col := &wallet.Collection{ID: "col-001", Type: "credential", Name: "我的证书", Owner: user1DID}
	if err := user1.AddCollection(col); err != nil {
		t.Fatalf("AddCollection failed: %v", err)
	}
	colGot, err := user1.GetCollection("col-001")
	if err != nil || colGot.ID != "col-001" {
		t.Fatalf("GetCollection failed: %v", err)
	}
	if err := user1.DeleteCollection("col-001"); err != nil {
		t.Fatalf("DeleteCollection failed: %v", err)
	}
	if _, err := user1.GetCollection("col-001"); err == nil {
		t.Fatalf("Collection should be deleted")
	}

	// 多用户隔离
	if err := user2.AddCredential(&wallet.Credential{ID: "vc-002", Issuer: user2DID}); err != nil {
		t.Fatalf("user2 AddCredential failed: %v", err)
	}
	if _, err := user1.GetCredential("vc-002"); err == nil {
		t.Fatalf("user1 should not see user2's credential")
	}
}

func TestWalletUserDIDDocumentKeyManagement(t *testing.T) {
	cfg := config.NewConfig()
	cfg.HuaweiCloudAccessKey = "dummy"
	cfg.DefaultAlgorithm = "ECDSA"

	// 生成密钥对
	keyPair, err := crypto.GenerateKeyPair(cfg, "ECDSA", "test-key")
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}
	// 计算DID标识符
	didIdentifier, err := did.CalculateDIDIdentifier(keyPair, "did:sbp:")
	if err != nil {
		t.Fatalf("CalculateDIDIdentifier failed: %v", err)
	}
	// 组装DID文档
	didDoc, err := did.AssembleDIDDocument(cfg, keyPair, "ECDSA", didIdentifier, nil)
	if err != nil {
		t.Fatalf("AssembleDIDDocument failed: %v", err)
	}
	user := &wallet.WalletUser{DID: didIdentifier, DIDDoc: didDoc, KeyManager: crypto.NewLocalKeyManager()}

	// 通过KeyManager生成新密钥
	keyID, _, err := user.KeyManager.Create(crypto.ECDSAP256)
	if err != nil {
		t.Fatalf("KeyManager.Create failed: %v", err)
	}
	// 用工具函数组装VerificationMethod
	newKey, err := did.NewVerificationMethodFromKeyManager(didIdentifier, keyID, "ECDSA", user.KeyManager)
	if err != nil {
		t.Fatalf("NewVerificationMethodFromKeyManager failed: %v", err)
	}
	if err := user.AddDIDKey(*newKey, "authentication"); err != nil {
		t.Fatalf("AddDIDKey failed: %v", err)
	}
	if len(user.DIDDoc.VerificationMethod) != 2 {
		t.Fatalf("Expected 2 keys, got %d", len(user.DIDDoc.VerificationMethod))
	}
	found := false
	for _, id := range user.DIDDoc.Authentication {
		if id == newKey.ID {
			found = true
		}
	}
	if !found {
		t.Fatalf("New key ID not found in authentication usage")
	}

	// 删除密钥
	if err := user.RemoveDIDKey(newKey.ID); err != nil {
		t.Fatalf("RemoveDIDKey failed: %v", err)
	}
	if len(user.DIDDoc.VerificationMethod) != 1 {
		t.Fatalf("Expected 1 key after removal, got %d", len(user.DIDDoc.VerificationMethod))
	}
	for _, id := range user.DIDDoc.Authentication {
		if id == newKey.ID {
			t.Fatalf("Key ID should be removed from authentication usage")
		}
	}
}

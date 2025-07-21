package main

import (
	"fmt"
	"github.com/helailiang/sbp-did-sdk-go/pkg/crypto"
	"github.com/helailiang/sbp-did-sdk-go/pkg/wallet"
)

func main() {
	// 1. 创建钱包实例
	w := wallet.NewWallet()

	// 2. 假设有两个用户DID
	user1DID := "did:sbp:user1"
	user2DID := "did:sbp:user2"

	// 3. 为每个用户选择不同的KeyManager（这里只演示本地，实际可用华为KMS）
	user1KM := crypto.NewLocalKeyManager()
	user2KM := crypto.NewLocalKeyManager() // 也可以用KMS实现

	// 4. 添加用户到钱包
	_ = w.AddUser(user1DID, user1KM)
	_ = w.AddUser(user2DID, user2KM)

	// 5. 获取用户对象
	user1, _ := w.GetUser(user1DID)
	user2, _ := w.GetUser(user2DID)

	// 6. 用户1创建并添加一个VC
	vc := &wallet.Credential{
		ID:      "vc-001",
		Context: []string{"https://www.w3.org/2018/credentials/v1"},
		Type:    []string{"VerifiableCredential"},
		Issuer:  user1DID,
		IssuanceDate: "2024-06-01T00:00:00Z",
		CredentialSubject: map[string]interface{}{"name": "张三", "age": 18},
	}
	_ = user1.AddCredential(vc)

	// 7. 用户1创建并添加一个密钥元数据
	keyID, pub, _ := user1KM.Create(crypto.ECDSAP256)
	key := &wallet.Key{
		ID:        keyID,
		Type:      "ECDSA",
		Controller: user1DID,
		PublicKey:  pub,
		Created:    "2024-06-01T00:00:00Z",
	}
	_ = user1.AddKey(key)

	// 8. 用户1创建并添加一个Collection
	col := &wallet.Collection{
		ID:   "col-001",
		Type: "credential",
		Name: "我的证书",
		Owner: user1DID,
	}
	_ = user1.AddCollection(col)

	// 9. 查询和打印
	vcGot, _ := user1.GetCredential("vc-001")
	fmt.Printf("用户1的VC: %+v\n", vcGot)
	keyGot, _ := user1.GetKey(keyID)
	fmt.Printf("用户1的Key: %+v\n", keyGot)
	colGot, _ := user1.GetCollection("col-001")
	fmt.Printf("用户1的Collection: %+v\n", colGot)

	// 10. 用户2也可以独立管理自己的VC、Key、Collection
	// ...（略）
} 
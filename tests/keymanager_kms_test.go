package tests

import (
	"os"
	"testing"
	"github.com/helailiang/sbp-did-sdk-go/pkg/crypto"
)

func TestHuaweiKMSKeyManager(t *testing.T) {
	endpoint := os.Getenv("KMS_ENDPOINT")
	ak := os.Getenv("KMS_AK")
	sk := os.Getenv("KMS_SK")
	projectId := os.Getenv("KMS_PROJECT_ID")
	km, err := crypto.NewHuaweiKMSKeyManager(endpoint, ak, sk, projectId)
	if err != nil {
		t.Fatal(err)
	}
	keyName := "test-key"
	_, err = km.GenerateKey("RSA", keyName)
	if err != nil {
		t.Fatal(err)
	}
	pub, err := km.GetPublicKey(keyName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PublicKey: %s", pub)
	sig, err := km.Sign(keyName, []byte("hello world"))
	if err != nil {
		t.Fatal(err)
	}
	valid, err := km.Verify(keyName, []byte("hello world"), sig)
	if err != nil {
		t.Fatal(err)
	}
	if !valid {
		t.Fatal("Signature verify failed")
	}
} 
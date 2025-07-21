package tests

import (
	"github.com/helailiang/sbp-did-sdk-go/pkg/crypto"
	"testing"
)

func TestLocalKeyManager(t *testing.T) {
	km := crypto.NewLocalKeyManager()
	keyName := "local-key"
	_, err := km.Create(crypto.ED25519, keyName)
	if err != nil {
		t.Fatal(err)
	}
	pub, err := km.GetPublicKey(keyName)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("PublicKey: %x", pub)
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

package rsautils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func TestKeySerialization(t *testing.T) {
	// generate a rsa key pair
	expectedPrivateKey, expectedPublicKey, err := GenerateRSAKeyPair(2048)
	if err != nil {
		t.Fatal(err)
	}
	// serialize the public key
	b64publicKey := B64EncodePublicKey(expectedPublicKey)
	// deserialize the public key
	actualPublicKey, err := PublicKeyFromBase64(b64publicKey)
	if err != nil {
		t.Fatal(err)
	}
	if !actualPublicKey.Equal(expectedPublicKey) {
		t.Fatalf("expected public key %v, got %v", expectedPublicKey, actualPublicKey)
	}

	// serialize the private key
	b64privateKey := B64EncodePrivateKey(expectedPrivateKey)
	// deserialize the private key
	actualPrivateKey, err := PrivateKeyFromBase64(b64privateKey)
	if err != nil {
		t.Fatal(err)
	}
	if !actualPrivateKey.Equal(expectedPrivateKey) {
		t.Fatalf("expected private key %v, got %v", expectedPrivateKey, actualPrivateKey)
	}

	// verify data signed with the private key can be verified with the public key
	data := []byte("hello world")
	hashed := crypto.SHA256.New()
	hashed.Write(data)
	hash := hashed.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, actualPrivateKey, crypto.SHA256, hash)
	if err != nil {
		t.Fatal(err)
	}
	err = rsa.VerifyPKCS1v15(actualPublicKey, crypto.SHA256, hash, signature)
	if err != nil {
		t.Fatal(err)
	}
}

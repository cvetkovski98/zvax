package service

import (
	"context"
	"testing"

	keys "github.com/cvetkovski98/zvax-keys/internal"
	"github.com/cvetkovski98/zvax-keys/internal/dto"
	"github.com/cvetkovski98/zvax-keys/internal/repository"
	rsautils "github.com/cvetkovski98/zvax-keys/internal/utils/rsa"
)

func GetMockedRepository(t *testing.T) keys.Repository {
	t.Helper()
	return repository.NewMockKeyRepository()
}

func GetService(t *testing.T) keys.Service {
	t.Helper()
	return NewKeyServiceImpl(GetMockedRepository(t))
}

func TestRegisterKey(t *testing.T) {
	bits := 2048
	_, publicKey, err := rsautils.GenerateRSAKeyPair(bits)
	if err != nil {
		t.Fatalf("failed to generate rsa key pair: %v", err)
	}
	b64PublicKey := rsautils.B64EncodePublicKey(publicKey)
	req := &dto.RegisterKey{
		Holder:      "oliver.cvetkovski@students.finki.ukim.mk",
		Affiliation: "students",
		Value:       b64PublicKey,
	}
	s := GetService(t)
	ctx := context.Background()
	key, certBytes, err := s.RegisterKey(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	if key == nil {
		t.Fatal("key is nil")
	}
	if certBytes == "" {
		t.Fatal("certBytes is empty")
	}
}

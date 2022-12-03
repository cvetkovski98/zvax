package rsautils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
)

func GenerateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

func B64EncodePrivateKey(privateKey *rsa.PrivateKey) string {
	privateKeyDER := x509.MarshalPKCS1PrivateKey(privateKey)
	b64privateKey := base64.StdEncoding.EncodeToString(privateKeyDER)
	return b64privateKey
}

func PrivateKeyFromBase64(b64privateKey string) (*rsa.PrivateKey, error) {
	privateKeyDER, err := base64.StdEncoding.DecodeString(b64privateKey)
	if err != nil {
		return nil, err
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyDER)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func B64EncodePublicKey(publicKey *rsa.PublicKey) string {
	publicKeyDER := x509.MarshalPKCS1PublicKey(publicKey)
	b64publicKey := base64.StdEncoding.EncodeToString(publicKeyDER)
	return b64publicKey
}

func PublicKeyFromBase64(b64publicKey string) (*rsa.PublicKey, error) {
	publicKeyDER, err := base64.StdEncoding.DecodeString(b64publicKey)
	if err != nil {
		return nil, err
	}
	publicKey, err := x509.ParsePKCS1PublicKey(publicKeyDER)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}

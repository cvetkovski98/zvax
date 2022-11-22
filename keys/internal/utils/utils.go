package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/pem"
	"github.com/pkg/errors"
	"math/big"
	"time"
)

func LoadCaKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(CaPrivateKeyPem))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "invalid private key")
	}
	return key, nil
}

func LoadCaCert() (*x509.Certificate, error) {
	privateKey, err := LoadCaKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load ca key")
	}
	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			CommonName:   "Zvax CA",
			Organization: []string{"Zvax"},
			Country:      []string{"MK"},
		},
		NotBefore:             time.Now().Add(-10 * time.Second),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}
	raw, err := x509.CreateCertificate(
		rand.Reader,
		cert,
		cert,
		&privateKey.PublicKey,
		privateKey,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create ca cert")
	}
	cert, err = x509.ParseCertificate(raw)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse ca cert")
	}
	return cert, nil
}

func ParseBase64PrivateKey(b64str string) (*rsa.PrivateKey, error) {
	raw, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		return nil, errors.Wrap(err, "invalid base64 encoded public key")
	}
	block, _ := pem.Decode(raw)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key")
	}
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "invalid private key")
	}
	return key, nil
}

func ParseBase64PublicKey(b64str string) (*rsa.PublicKey, error) {
	rawB64, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		return nil, errors.Wrap(err, "invalid base64 encoded public key")
	}
	block, _ := pem.Decode(rawB64)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("pem is not a public key")
	}
	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse PEM into public key")
	}
	rsaKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("public key is not an RSA key")
	}
	return rsaKey, nil
}

package service

import (
	"context"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"math/big"
	"time"

	keys "github.com/cvetkovski98/zvax-keys/internal"
	"github.com/cvetkovski98/zvax-keys/internal/dto"
	"github.com/cvetkovski98/zvax-keys/internal/mapper"
	cautils "github.com/cvetkovski98/zvax-keys/internal/utils/ca"
	rsautils "github.com/cvetkovski98/zvax-keys/internal/utils/rsa"
	"github.com/pkg/errors"
)

type impl struct {
	kr keys.Repository
}

// signCertificate generates a signed certificate for a template using the CA Certificate and CA Private Key.
// returns the bytes of the PEM-encoded certificate.
func (s *impl) signCertificate(template *x509.Certificate) ([]byte, error) {
	rootKey, err := cautils.LoadCaKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load ca key")
	}
	rootCert, err := cautils.LoadCaCert()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load ca cert")
	}
	raw, err := x509.CreateCertificate(
		rand.Reader,
		template,
		rootCert,
		template.PublicKey,
		rootKey,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create certificate")
	}
	pemBytes := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: raw,
	})
	return pemBytes, nil
}

func (s *impl) RegisterKey(ctx context.Context, key *dto.RegisterKey) (*dto.Key, string, error) {
	publicKey, err := rsautils.PublicKeyFromBase64(key.Value)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to parse base64 into public key")
	}
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(1<<32))
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to generate serial number")
	}
	emailAddressOid := asn1.ObjectIdentifier{1, 2, 840, 113549, 1, 9, 1}
	cert := x509.Certificate{
		SerialNumber: serialNumber,
		PublicKey:    publicKey,
		Subject: pkix.Name{
			CommonName:   "zvax",
			Organization: []string{key.Affiliation},
			Names: []pkix.AttributeTypeAndValue{
				{Type: emailAddressOid, Value: key.Holder},
			},
		},
		NotBefore:   time.Now().Add(-time.Second * 10),
		NotAfter:    time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:    x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		IsCA:        false,
	}
	certPemBytes, err := s.signCertificate(&cert)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to create certificate")
	}
	certBase64 := base64.StdEncoding.EncodeToString(certPemBytes)
	keyModel := mapper.RegisterKeyDtoToModel(key)
	created, err := s.kr.InsertOne(ctx, keyModel)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to insert key")
	}
	return mapper.KeyModelToDto(created), certBase64, nil
}

func (s *impl) ListKeys(ctx context.Context, holder string) (*dto.Keys, error) {
	keys, err := s.kr.FindAllByHolder(ctx, holder)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find keys")
	}
	return mapper.KeysModelToDto(keys), nil

}

func (s *impl) GetKey(ctx context.Context, id int) (*dto.Key, error) {
	key, err := s.kr.FindOneById(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find key")
	}
	return mapper.KeyModelToDto(key), nil
}

func NewKeyServiceImpl(userRepository keys.Repository) keys.Service {
	return &impl{
		kr: userRepository,
	}
}

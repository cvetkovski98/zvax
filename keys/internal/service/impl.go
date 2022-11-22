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

	"github.com/cvetkovski98/zvax-keys/internal"
	"github.com/cvetkovski98/zvax-keys/internal/model"
	"github.com/cvetkovski98/zvax-keys/internal/model/dto"
	"github.com/cvetkovski98/zvax-keys/internal/utils"
	"github.com/pkg/errors"
)

type keyServiceImpl struct {
	keyRepository internal.KeyRepository
}

// signCertificate generates a signed certificate for a template using the CA Certificate and CA Private Key.
// returns the bytes of the PEM-encoded certificate.
func (service *keyServiceImpl) signCertificate(template *x509.Certificate) ([]byte, error) {
	rootKey, err := utils.LoadCaKey()
	if err != nil {
		return nil, errors.Wrap(err, "failed to load ca key")
	}
	rootCert, err := utils.LoadCaCert()
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

func (service *keyServiceImpl) RegisterKey(ctx context.Context, key *dto.RegisterKeyInDto) (*model.Key, *string, error) {
	publicKey, err := utils.ParseBase64PublicKey(key.PublicKey)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to parse base64 into public key")
	}
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(1<<32))
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate serial number")
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
	certPemBytes, err := service.signCertificate(&cert)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create certificate")
	}
	certBase64 := base64.StdEncoding.EncodeToString(certPemBytes)
	keyIn := &model.Key{
		Holder:      key.Holder,
		Affiliation: key.Affiliation,
		Value:       key.PublicKey,
	}
	created, err := service.keyRepository.InsertOne(ctx, keyIn)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to insert key")
	}
	return created, &certBase64, err
}

func (service *keyServiceImpl) ListKeys(ctx context.Context) ([]*model.Key, error) {
	return service.keyRepository.FindAll(ctx)
}

func (service *keyServiceImpl) GetKey(ctx context.Context, keyId int64) (*model.Key, error) {
	return service.keyRepository.FindOneById(ctx, keyId)
}

func NewKeyServiceImpl(userRepository internal.KeyRepository) internal.KeyService {
	return &keyServiceImpl{
		keyRepository: userRepository,
	}
}

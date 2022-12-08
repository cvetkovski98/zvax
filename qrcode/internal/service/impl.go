package service

import (
	"context"
	"errors"

	qrcode "github.com/cvetkovski98/zvax/zvax-qrcode/internal"
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/dto"
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/mapper"
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/model"
	errutil "github.com/cvetkovski98/zvax/zvax-qrcode/internal/utils/error"
	qrutil "github.com/cvetkovski98/zvax/zvax-qrcode/internal/utils/qr"
)

type impl struct {
	r qrcode.Repository
}

func (s *impl) validateCreateQRCode(ctx context.Context, request *dto.CreateQRCode) error {
	emailEmpty := request.Email == nil || *request.Email == ""

	if request.Store && emailEmpty {
		return errors.New("email is required when storing QR code")
	}

	if !emailEmpty {
		qr, err := s.r.FindOneByEmail(ctx, *request.Email)
		err = errutil.ParseError(err)
		if err != nil {
			return err
		}
		if qr != nil {
			return errors.New("QR code already exists for this email")
		}
	}
	return nil
}

func (s *impl) CreateQRCode(ctx context.Context, request *dto.CreateQRCode) (*dto.QR, error) {
	if err := s.validateCreateQRCode(ctx, request); err != nil {
		return nil, err
	}

	content, err := qrutil.GenerateQRCode(request.Content)
	if err != nil {
		return nil, err
	}

	result := &dto.QR{
		Content: content,
	}

	if request.Store {
		QRin := &model.QR{
			Email:   *request.Email,
			Content: content,
		}
		_, err := s.r.InsertOne(ctx, QRin)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func (s *impl) GetQRCode(ctx context.Context, request *dto.GetQRCode) (*dto.StoredQR, error) {
	result, err := s.r.FindOneByEmail(ctx, request.Email)
	if err != nil {
		return nil, err
	}
	return mapper.QRModelToDto(result), nil
}

func NewQRCodeService(r qrcode.Repository) qrcode.Service {
	return &impl{r: r}
}

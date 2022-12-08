package qrcode

import (
	"context"

	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/dto"
)

type Service interface {
	CreateQRCode(context.Context, *dto.CreateQRCode) (*dto.QR, error)
	GetQRCode(context.Context, *dto.GetQRCode) (*dto.StoredQR, error)
}

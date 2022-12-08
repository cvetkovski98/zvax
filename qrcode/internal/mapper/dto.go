package mapper

import (
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/dto"
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/model"
)

func QRModelToDto(qr *model.QR) *dto.StoredQR {
	return &dto.StoredQR{
		ID:        qr.ID,
		Email:     qr.Email,
		Content:   qr.Content,
		CreatedAt: qr.CreatedAt,
		UpdatedAt: qr.UpdatedAt,
	}
}

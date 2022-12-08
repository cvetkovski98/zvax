package mapper

import (
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/dto"
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/model"
)

func QRCodeDtoToModel(qr *dto.StoredQR) *model.QR {
	return &model.QR{
		ID:        qr.ID,
		Email:     qr.Email,
		Content:   qr.Content,
		CreatedAt: qr.CreatedAt,
		UpdatedAt: qr.UpdatedAt,
	}
}

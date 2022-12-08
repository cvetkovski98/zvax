package delivery

import (
	"github.com/cvetkovski98/zvax-common/gen/pbqr"
	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/dto"
)

func CreateQRCodeRequestToDto(req *pbqr.CreateQRCodeRequest) *dto.CreateQRCode {
	return &dto.CreateQRCode{
		Email:   req.Email,
		Content: req.Content,
		Store:   req.Store,
	}
}

func GetQRCodeRequestToDto(req *pbqr.GetQRCodeRequest) *dto.GetQRCode {
	return &dto.GetQRCode{
		Email: req.Email,
	}
}

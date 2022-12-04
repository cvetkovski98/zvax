package service

import (
	"bytes"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	qrcode "github.com/cvetkovski98/zvax/zvax-qrcode/internal"
)

type impl struct {
}

func (*impl) Generate(payload string) ([]byte, error) {
	qrCode, err := qr.Encode(payload, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}
	qrCode, err = barcode.Scale(qrCode, 256, 256)
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer
	png.Encode(&buffer, qrCode)
	return buffer.Bytes(), err
}

func NewQRCodeService() qrcode.Service {
	return &impl{}
}

package qrcode

import (
	"context"

	"github.com/cvetkovski98/zvax/zvax-qrcode/internal/model"
)

type Repository interface {
	InsertOne(context.Context, *model.QR) (*model.QR, error)
	FindOneByEmail(context.Context, string) (*model.QR, error)
}

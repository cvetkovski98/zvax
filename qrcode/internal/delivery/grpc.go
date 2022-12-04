package delivery

import (
	"context"

	"github.com/cvetkovski98/zvax-common/gen/pbqr"
	qrcode "github.com/cvetkovski98/zvax/zvax-qrcode/internal"
)

type server struct {
	s qrcode.Service

	pbqr.UnimplementedQRServer
}

func (s *server) GenerateQR(ctx context.Context, request *pbqr.QRRequest) (*pbqr.QRResponse, error) {
	content, err := s.s.Generate(request.Content)
	if err != nil {
		return nil, err
	}
	response := &pbqr.QRResponse{
		Qr: content,
	}
	return response, nil
}

func NewQRServer(s qrcode.Service) pbqr.QRServer {
	return &server{
		s: s,
	}
}

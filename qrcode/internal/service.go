package qrcode

type Service interface {
	Generate(payload string) ([]byte, error)
}

package dto

import "time"

type CreateQRCode struct {
	Email   *string
	Content string
	Store   bool
}

type GetQRCode struct {
	Email string
}

type QR struct {
	Content []byte
}

type StoredQR struct {
	ID        int
	Email     string
	Content   []byte
	CreatedAt time.Time
	UpdatedAt time.Time
}

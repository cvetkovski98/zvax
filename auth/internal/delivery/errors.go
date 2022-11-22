package delivery

import "errors"

var (
	ErrMissingMetadata = errors.New("missing metadata")
	ErrMissingCookies  = errors.New("missing cookies")
	ErrMissingToken    = errors.New("missing token")
)

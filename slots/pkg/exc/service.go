package exc

import "errors"

var ErrRequiredFieldsMissing = errors.New("required fields must not be nil")

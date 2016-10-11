package depl

import (
	"errors"
)

var (
	ErrEmptyString  = errors.New("Empty string")
	ErrInvalidEmail = errors.New("Invalid email address")
)

package core

import (
	"errors"
)

var (
	ErrInvalidSize      error = errors.New("invalid size, must be greater than 0")
	ErrInvalidDirection error = errors.New("invalid direction, must be one of: up, down, left, right")
)

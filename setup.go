package blic

import (
	"errors"
)

var (
	ErrInvalidLink    = errors.New("Invalid link")
	ErrOrphanLocation = errors.New("Location has no neighbours")
)

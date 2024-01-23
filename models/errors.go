package models

import "errors"

var (
	ErrEmailTaken = errors.New("models: email address is already in use")
	ErrNotFound   = errors.New("models: resource could not be found.")
)

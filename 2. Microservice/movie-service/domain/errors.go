package domain

import "errors"

var (
	ErrParamTitleMinCharacter = errors.New("param title must at least 3 characters")
	ErrInvalidParam           = errors.New("invalid param")
)

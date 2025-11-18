package errorapp

import "errors"

var (
	ErrInitConfig   = errors.New("error init config")
	ErrCreateServer = errors.New("error create openapi server")
)

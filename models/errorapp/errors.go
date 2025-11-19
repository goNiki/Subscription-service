package errorapp

import "errors"

var (
	ErrInitConfig         = errors.New("error init config")
	ErrCreateServer       = errors.New("error create openapi server")
	ErrParceConfigDB      = errors.New("failed to parce config DB")
	ErrCreateConnectionDB = errors.New("failed to create connection pool")
	ErrPingDB             = errors.New("failed to ping database")
)

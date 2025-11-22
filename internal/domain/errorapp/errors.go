package errorapp

import "errors"

var (
	ErrInitConfig           = errors.New("error init config")
	ErrCreateServer         = errors.New("error create openapi server")
	ErrParceConfigDB        = errors.New("failed to parce config DB")
	ErrCreateConnectionDB   = errors.New("failed to create connection pool")
	ErrPingDB               = errors.New("failed to ping database")
	ErrParceDate            = errors.New("invalid date format, expected MM-YYYY")
	ErrValidation           = errors.New("invalid date")
	ErrConverter            = errors.New("converter error")
	ErrDBInternal           = errors.New("database internal error")
	ErrCreateSubscription   = errors.New("error create subscription")
	ErrNotFoundSubscription = errors.New("subscriptopn not found")
	ErrInvalidUUID          = errors.New("invalid uuid")
)

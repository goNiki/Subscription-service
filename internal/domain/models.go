package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID
	ServiceName string
	Price       int64
	UserID      uuid.UUID
	StartDate   time.Time
	EndDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ParamsGetCost struct {
	UserID      *uuid.UUID
	ServiceName *string
	StartDate   time.Time
	EndDate     time.Time
}

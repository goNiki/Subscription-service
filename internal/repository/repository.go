package repository

import (
	"context"

	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/google/uuid"
)

type SubscriptionsRepository interface {
	CreateSubscriptions(ctx context.Context, sub *domain.Subscription) (*domain.Subscription, error)
	GetSubscription(ctx context.Context, id uuid.UUID) (*domain.Subscription, error)
	DeleteSubscription(ctx context.Context, id uuid.UUID) error
	UpdateSubscription(ctx context.Context, id uuid.UUID, sub *domain.Subscription) (*domain.Subscription, error)
	Getsubscriptions(ctx context.Context, args []any, conditions []string) ([]*domain.Subscription, error)
	GetTotalCost(ctx context.Context, args []any, conditions []string) (int64, error)
}

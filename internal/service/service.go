package service

import (
	"context"

	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/google/uuid"
)

type SubscriptionService interface {
	CreateSubscription(ctx context.Context, subscription *domain.Subscription) (*domain.Subscription, error)
	GetSubscription(ctx context.Context, id uuid.UUID) (*domain.Subscription, error)
	DeleteSubscription(ctx context.Context, id uuid.UUID) error
	UpdateSubscription(ctx context.Context, id uuid.UUID, sub *domain.Subscription) (*domain.Subscription, error)
	GetSubscriptions(ctx context.Context, service_name *string, userId *uuid.UUID) ([]*domain.Subscription, error)
	GetTotalCostSubscriptions(ctx context.Context, params domain.ParamsGetCost) (int64, error)
}

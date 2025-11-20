package repository

import "context"

type SubscriptionsRepository interface {
	CreateSubscriptions(ctx context.Context) error
}

package subscrtiptionsservice

import (
	"github.com/goNiki/Subscription-service/internal/repository"
)

type Service struct {
	SubscriptionRepo repository.SubscriptionsRepository
}

func NewSubscriptionsService(subscriptionsRepo repository.SubscriptionsRepository) *Service {
	return &Service{
		SubscriptionRepo: subscriptionsRepo,
	}
}

package subscrtiptionsservice

import (
	"context"
	"fmt"

	"github.com/goNiki/Subscription-service/internal/domain"
)

func (s *Service) CreateSubscription(ctx context.Context, subscription *domain.Subscription) (*domain.Subscription, error) {

	sub, err := s.SubscriptionRepo.CreateSubscriptions(ctx, subscription)
	if err != nil {
		return &domain.Subscription{}, fmt.Errorf("error create subscription: %w", err)
	}

	return sub, nil
}

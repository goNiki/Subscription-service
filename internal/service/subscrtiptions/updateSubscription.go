package subscrtiptionsservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func (s *Service) UpdateSubscription(ctx context.Context, id uuid.UUID, sub *domain.Subscription) (*domain.Subscription, error) {

	subscriptions, err := s.SubscriptionRepo.UpdateSubscription(ctx, id, sub)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &domain.Subscription{}, fmt.Errorf("%w: %v", errorapp.ErrNotFoundSubscription, err)
		}
		return &domain.Subscription{}, fmt.Errorf("%w: %v", errorapp.ErrDBInternal, err)
	}

	return subscriptions, nil

}

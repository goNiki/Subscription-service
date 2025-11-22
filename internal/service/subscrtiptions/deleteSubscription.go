package subscrtiptionsservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	"github.com/google/uuid"
)

func (s *Service) DeleteSubscription(ctx context.Context, id uuid.UUID) error {
	err := s.SubscriptionRepo.DeleteSubscription(ctx, id)
	if err != nil {
		if errors.Is(err, errorapp.ErrNotFoundSubscription) {
			return err
		}
		return fmt.Errorf("%w: %v", errorapp.ErrDBInternal, err)
	}
	return nil
}

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

func (s *Service) GetSubscriptions(ctx context.Context, service_name *string, userId *uuid.UUID) ([]*domain.Subscription, error) {

	var args []any
	var conditions []string

	if service_name != nil {
		conditions = append(conditions, "service_name = $"+fmt.Sprint(len(args)+1))
		args = append(args, *service_name)
	}

	if userId != nil {
		conditions = append(conditions, "user_id = $"+fmt.Sprint(len(args)+1))
		args = append(args, *userId)
	}

	resp, err := s.SubscriptionRepo.Getsubscriptions(ctx, args, conditions)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%w: %v", errorapp.ErrNotFoundSubscription, err)
		}
		return nil, fmt.Errorf("%w: %v", errorapp.ErrDBInternal, err)
	}
	return resp, nil

}

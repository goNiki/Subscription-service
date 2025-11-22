package subscrtiptionsservice

import (
	"context"
	"errors"
	"fmt"

	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	"github.com/jackc/pgx/v5"
)

func (s *Service) GetTotalCostSubscriptions(ctx context.Context, params domain.ParamsGetCost) (int64, error) {
	var args []any
	var condition []string

	if params.UserID != nil {
		condition = append(condition, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, params.UserID)
	}

	if params.ServiceName != nil {
		condition = append(condition, fmt.Sprintf("service_name = $%d", len(args)+1))
		args = append(args, params.ServiceName)
	}

	if !params.StartDate.IsZero() {
		condition = append(condition, fmt.Sprintf("start_date <= $%d", len(args)+1))
		args = append(args, params.StartDate)
	}

	if !params.EndDate.IsZero() {
		condition = append(condition, fmt.Sprintf("(end_date >= $%d OR end_date IS NULL)", len(args)+1))
		args = append(args, params.EndDate)
	}

	totalCost, err := s.SubscriptionRepo.GetTotalCost(ctx, args, condition)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, fmt.Errorf("%w: %v", errorapp.ErrNotFoundSubscription, err)
		}
		return 0, fmt.Errorf("%w: %v", errorapp.ErrDBInternal, err)
	}

	return totalCost, nil
}

package subscriptionsrepo

import (
	"context"
	"fmt"
	"strings"

	"github.com/goNiki/Subscription-service/internal/domain"
)

func (r *Repository) Getsubscriptions(ctx context.Context, args []any, conditions []string) ([]*domain.Subscription, error) {
	const op = "repository.subscriptions.getsubscriptions"

	query := `SELECT id, service_name, price, user_id, start_date, end_date, created_at, updated_at FROM subscriptions`

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	rows, err := r.Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var result []*domain.Subscription

	for rows.Next() {
		sub := &domain.Subscription{}
		err := rows.Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate, &sub.CreatedAt, &sub.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		result = append(result, sub)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return result, nil

}

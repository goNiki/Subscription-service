package subscriptionsrepo

import (
	"context"
	"fmt"
	"strings"
)

func (r *Repository) GetTotalCost(ctx context.Context, args []any, conditions []string) (int64, error) {
	const op = "repository.subscriptions.getTotalCost"

	query := `SELECT COALESCE(SUM(price), 0) as total_cost
			FROM subscriptions`

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	var totalCost int64

	err := r.Pool.QueryRow(ctx, query, args...).Scan(&totalCost)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return totalCost, nil

}

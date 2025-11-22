package subscriptionsrepo

import (
	"context"
	"fmt"

	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/google/uuid"
)

func (r *Repository) GetSubscription(ctx context.Context, id uuid.UUID) (*domain.Subscription, error) {
	const op = "repository.subscriptions.getsubscription"

	query := `SELECT id, service_name, price, user_id, start_date, end_date, created_at, updated_at FROM subscriptions WHERE id = $1`

	var sub domain.Subscription

	err := r.Pool.QueryRow(ctx, query, id).Scan(&sub.ID, &sub.ServiceName, &sub.Price, &sub.UserID, &sub.StartDate, &sub.EndDate, &sub.CreatedAt, &sub.UpdatedAt)
	if err != nil {
		return &domain.Subscription{}, fmt.Errorf("%s: %w", op, err)
	}
	return &sub, nil
}

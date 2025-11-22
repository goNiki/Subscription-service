package subscriptionsrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/goNiki/Subscription-service/internal/repository/converter"
	"github.com/google/uuid"
)

func (r *Repository) UpdateSubscription(ctx context.Context, id uuid.UUID, sub *domain.Subscription) (*domain.Subscription, error) {
	const op = "repository.subscriptions.updatesubscription"

	dbsub := converter.SubscriptionToModel(sub)

	query := `UPDATE subscriptions 
			SET service_name = $2,
				price = $3,
				user_id = $4,
				start_date = $5,
				end_date = $6,
				updated_at = $7 
			WHERE id = $1
			RETURNING created_at`

	updateAt := time.Now()
	var createdAt time.Time

	err := r.Pool.QueryRow(ctx, query, id, dbsub.ServiceName, dbsub.Price, dbsub.UserID, dbsub.StartDate, dbsub.EndDate, updateAt).Scan(&createdAt)
	if err != nil {
		return &domain.Subscription{}, fmt.Errorf("%s: %w", op, err)
	}

	resp := domain.Subscription{
		ID:          id,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updateAt,
	}

	return &resp, nil
}

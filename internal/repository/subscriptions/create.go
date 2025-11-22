package subscriptionsrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/goNiki/Subscription-service/internal/repository/converter"
	"github.com/google/uuid"
)

func (r *Repository) CreateSubscriptions(ctx context.Context, sub *domain.Subscription) (*domain.Subscription, error) {
	const op = "repository.subscriptions.createsubscription"

	dbsub := converter.SubscriptionToModel(sub)

	query := `INSERT INTO subscriptions (id, service_name, price, user_id, start_date, end_date, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	          RETURNING id, created_at, updated_at`

	id := uuid.New()
	now := time.Now()

	var ID uuid.UUID
	var createdAt, updatedAt time.Time

	err := r.Pool.QueryRow(ctx, query,
		id, dbsub.ServiceName, dbsub.Price, dbsub.UserID,
		dbsub.StartDate, dbsub.EndDate, now, now,
	).Scan(&ID, &createdAt, &updatedAt)

	if err != nil {
		return &domain.Subscription{}, fmt.Errorf("%s: %w", op, err)
	}

	sub.ID = ID
	sub.CreatedAt = createdAt
	sub.UpdatedAt = updatedAt

	return sub, nil
}

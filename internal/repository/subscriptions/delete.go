package subscriptionsrepo

import (
	"context"
	"fmt"

	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	"github.com/google/uuid"
)

func (r *Repository) DeleteSubscription(ctx context.Context, id uuid.UUID) error {
	const op = "repository.subscriptions.deletesubscription"

	query := `DELETE FROM subscriptions WHERE id = $1`

	tag, err := r.Pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if tag.RowsAffected() == 0 {
		return fmt.Errorf("%s: %w", op, errorapp.ErrNotFoundSubscription)
	}

	return nil
}

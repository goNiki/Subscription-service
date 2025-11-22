package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	subV1 "github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
	"github.com/google/uuid"
)

func (a *Api) SubscriptionDeleteByID(ctx context.Context, params subV1.SubscriptionDeleteByIDParams) (subV1.SubscriptionDeleteByIDRes, error) {
	const op = "subscriptionDeleteByID"

	if len(params.SubUUID) != 36 {
		a.logError(ctx, op, errorapp.ErrInvalidUUID)
		return &subV1.BadRequestError{
			Code:    400,
			Message: "Неверно введен UUID подписки",
		}, nil
	}

	id := uuid.MustParse(params.SubUUID)

	err := a.subscriptionsService.DeleteSubscription(ctx, id)
	if err != nil {
		a.logError(ctx, op, err)
		switch {
		case errors.Is(err, errorapp.ErrNotFoundSubscription):
			return &subV1.NotFoundError{
				Code:    404,
				Message: fmt.Sprintf("Подписки с UUID %s нет", id),
			}, nil
		default:
			return &subV1.InternalServerError{
				Code:    500,
				Message: "Внутренняя ошибка сервера",
			}, nil
		}
	}

	return &subV1.SubscriptionDeleteByIDNoContent{}, nil
}

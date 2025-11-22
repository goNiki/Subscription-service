package v1

import (
	"context"
	"errors"
	"fmt"

	"github.com/goNiki/Subscription-service/internal/converter"
	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	subV1 "github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
	"github.com/google/uuid"
)

func (a *Api) SubscriptionGetByID(ctx context.Context, params subV1.SubscriptionGetByIDParams) (subV1.SubscriptionGetByIDRes, error) {
	const op = "SubscriptionGetByID"

	if len(params.SubUUID) != 36 {
		a.logError(ctx, op, errorapp.ErrInvalidUUID)
		return &subV1.BadRequestError{
			Code:    400,
			Message: "Неверно введен UUID подписки",
		}, nil
	}

	id := uuid.MustParse(params.SubUUID)

	sub, err := a.subscriptionsService.GetSubscription(ctx, id)
	if err != nil {
		a.logError(ctx, op, err)
		switch {
		case errors.Is(err, errorapp.ErrNotFoundSubscription):
			return &subV1.NotFoundError{
				Code:    404,
				Message: fmt.Sprintf("Подписки c id %s в базе нет", id),
			}, nil

		default:
			return &subV1.InternalServerError{
				Code:    500,
				Message: "Ошибка в системе",
			}, nil
		}
	}

	resp := converter.SubscriptionToDTO(sub)

	return &resp, nil
}

package v1

import (
	"context"
	"errors"

	"github.com/goNiki/Subscription-service/internal/converter"
	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	subV1 "github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
	"github.com/google/uuid"
)

func (a *Api) GetSubscription(ctx context.Context, req subV1.OptGetSubscriptionsRequest) (subV1.GetSubscriptionRes, error) {
	const op = "GetSubscription"

	var serviceName *string
	var userId *uuid.UUID

	if req.Set {
		if req.Value.ServiceName.Set {
			serviceName = &req.Value.ServiceName.Value
		}
		if req.Value.UserID.Set {
			i := req.Value.UserID.Value
			userId = &i
		}
	}

	subs, err := a.subscriptionsService.GetSubscriptions(ctx, serviceName, userId)
	if err != nil {
		a.logError(ctx, op, err)
		switch {
		case errors.Is(err, errorapp.ErrNotFoundSubscription):
			return &subV1.NotFoundError{
				Code:    404,
				Message: "Подписок с указанными параметрами нет",
			}, nil

		default:
			return &subV1.InternalServerError{
				Code:    500,
				Message: "Ошибка в системе",
			}, nil
		}
	}

	subsctiptions := converter.SubscriptionsSliasToDTO(subs)
	total := len(subsctiptions)

	return &subV1.GetSubscriptionsResponse{
		Subscriptions: subsctiptions,
		Total:         int64(total),
	}, nil
}

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

func (a *Api) SubscriptionUpdateByID(ctx context.Context, req *subV1.SubscriptionsReqDto, params subV1.SubscriptionUpdateByIDParams) (subV1.SubscriptionUpdateByIDRes, error) {
	const op = "SubscriptionUpdateByID"

	if len(params.SubUUID) != 36 {
		a.logError(ctx, op, errorapp.ErrInvalidUUID)
		return &subV1.BadRequestError{
			Code:    400,
			Message: "Неверно введен UUID подписки",
		}, nil
	}

	domain, err := converter.SubscriptionDtoToModel(req)
	if err != nil {
		a.logError(ctx, op, err)
		switch {
		case errors.Is(err, errorapp.ErrParceDate):
			return &subV1.BadRequestError{
				Code:    400,
				Message: "неверно указана дата",
			}, nil
		default:
			return &subV1.InternalServerError{
				Code:    500,
				Message: "Внутренняя ошибка сервера",
			}, nil
		}
	}

	id := uuid.MustParse(params.SubUUID)

	subscriptions, err := a.subscriptionsService.UpdateSubscription(ctx, id, domain)
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

	resp := converter.SubscriptionToDTO(subscriptions)

	return &resp, nil
}

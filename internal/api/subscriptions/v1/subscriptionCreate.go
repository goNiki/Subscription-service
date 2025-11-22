package v1

import (
	"context"
	"errors"

	"github.com/goNiki/Subscription-service/internal/converter"
	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	subV1 "github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
)

func (a *Api) SubscriptionCreate(ctx context.Context, req *subV1.SubscriptionsReqDto) (subV1.SubscriptionCreateRes, error) {
	const op = "SubscriptionCreate"

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

	date, err := a.subscriptionsService.CreateSubscription(ctx, domain)
	if err != nil {
		a.logError(ctx, op, err)
		switch {
		default:
			return &subV1.InternalServerError{
				Code:    500,
				Message: "Внутренняя ошибка сервера",
			}, nil
		}

	}

	resp := converter.SubscriptionToDTO(date)
	return &resp, nil
}

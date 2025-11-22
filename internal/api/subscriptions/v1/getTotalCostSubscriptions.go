package v1

import (
	"context"

	"github.com/goNiki/Subscription-service/internal/converter"
	subV1 "github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
)

func (a *Api) GetTotalCostSubscriptions(ctx context.Context, params subV1.GetTotalCostSubscriptionsParams) (subV1.GetTotalCostSubscriptionsRes, error) {
	const op = "GetTotalCostSubscriptions"

	domain, err := converter.ParamsGetTotalCostToModel(params.ServiceName, params.UserID, params.StartDate, params.EndDate)
	if err != nil {
		a.logError(ctx, op, err)
		return &subV1.BadRequestError{
			Code:    400,
			Message: "неверно указана дата",
		}, nil
	}

	totalCost, err := a.subscriptionsService.GetTotalCostSubscriptions(ctx, domain)
	if err != nil {
		a.logError(ctx, op, err)
		switch {
		default:
			return &subV1.InternalServerError{
				Code:    500,
				Message: "Ошибка в системе",
			}, nil
		}
	}

	resp := subV1.GetTotalcostResponse{
		TotalCost: totalCost,
		Period: subV1.Period{
			StartDate: params.StartDate,
			EndDate:   params.EndDate,
		},
		Filters: subV1.Filters{
			UserID:      subV1.NewOptString(params.UserID.Value.String()),
			ServiceName: params.ServiceName,
		},
	}

	return &resp, nil
}

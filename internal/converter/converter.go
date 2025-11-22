package converter

import (
	"fmt"
	"time"

	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/goNiki/Subscription-service/internal/domain/errorapp"
	subV1 "github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
)

func SubscriptionDtoToModel(sub *subV1.SubscriptionsReqDto) (*domain.Subscription, error) {
	startDate, endDate, err := DateToTime(sub.StartDate, sub.EndDate)
	if err != nil {
		return &domain.Subscription{}, err
	}
	return &domain.Subscription{
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   startDate,
		EndDate:     endDate,
	}, nil

}

func SubscriptionToDTO(sub *domain.Subscription) subV1.SubscriptionsRespDto {
	startDate, endDate := TimeToDate(sub.StartDate, sub.EndDate)
	return subV1.SubscriptionsRespDto{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   sub.CreatedAt,
		UpdatedAt:   sub.UpdatedAt,
	}
}

func SubscriptionsSliasToDTO(subs []*domain.Subscription) []subV1.SubscriptionsRespDto {
	var resp []subV1.SubscriptionsRespDto

	for _, v := range subs {
		resp = append(resp, SubscriptionToDTO(v))
	}

	return resp

}

func ParamsGetTotalCostToModel(serviceName subV1.OptString, userID subV1.OptUUID, startDate string, endDate string) (domain.ParamsGetCost, error) {
	var params domain.ParamsGetCost
	var err error
	if serviceName.Set == true {
		params.ServiceName = &serviceName.Value
	}

	if userID.Set == true {
		id := userID.Value
		params.UserID = &id
	}

	params.StartDate, err = ParceMonthYear(startDate)
	if err != nil {
		return domain.ParamsGetCost{}, err
	}

	params.EndDate, err = ParceMonthYear(endDate)
	if err != nil {
		return domain.ParamsGetCost{}, err
	}

	return params, nil
}

func DateToTime(start string, end subV1.OptString) (time.Time, *time.Time, error) {
	startDate, err := ParceMonthYear(start)
	if err != nil {
		return time.Time{}, &time.Time{}, err
	}

	if end.Set == false {
		return startDate, nil, nil
	}
	endDate, err := ParceMonthYear(end.Value)
	if err != nil {
		return time.Time{}, &time.Time{}, err
	}

	return startDate, &endDate, nil
}

func TimeToDate(start time.Time, end *time.Time) (string, subV1.OptString) {
	if end == nil {
		return start.Format("01-2006"), subV1.OptString{Set: false}
	}
	return start.Format("01-2006"), subV1.OptString{Value: end.Format("01-2006"), Set: true}
}

func ParceMonthYear(dataStr string) (time.Time, error) {
	t, err := time.Parse("01-2006", dataStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("%w: %v", errorapp.ErrParceDate, err)
	}
	return t, nil
}

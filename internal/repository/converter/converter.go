package converter

import (
	"github.com/goNiki/Subscription-service/internal/domain"
	"github.com/goNiki/Subscription-service/internal/repository/models"
)

func SubscriptionToModel(sub *domain.Subscription) models.Subscription {
	return models.Subscription{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
		CreatedAt:   sub.CreatedAt,
		UpdatedAt:   sub.UpdatedAt,
	}
}

func SubscriptionToDomain(sub *models.Subscription) domain.Subscription {
	return domain.Subscription{
		ID:          sub.ID,
		ServiceName: sub.ServiceName,
		Price:       sub.Price,
		UserID:      sub.UserID,
		StartDate:   sub.StartDate,
		EndDate:     sub.EndDate,
		CreatedAt:   sub.CreatedAt,
		UpdatedAt:   sub.UpdatedAt,
	}
}

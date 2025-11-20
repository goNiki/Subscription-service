package v1

import (
	"log/slog"

	"github.com/goNiki/Subscription-service/internal/service"
)

type Api struct {
	log                  *slog.Logger
	subscriptionsService service.SubscriptionService
}

func NewSubscriptionsApi(log *slog.Logger, SubscriptionsService service.SubscriptionService) *Api {
	return &Api{
		log:                  log,
		subscriptionsService: SubscriptionsService,
	}
}

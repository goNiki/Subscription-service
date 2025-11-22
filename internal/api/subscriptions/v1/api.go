package v1

import (
	"context"
	"log/slog"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/goNiki/Subscription-service/internal/infrastructure/logger/sl"
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

func (a *Api) logError(ctx context.Context, operation string, err error) {
	a.log.Error("operation Error",
		slog.String("operation", operation),
		slog.String("request_id", middleware.GetReqID(ctx)),
		sl.Error(err))
}

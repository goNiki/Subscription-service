package v1

import (
	"context"

	subV1 "github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
)

func (a *Api) NewError(ctx context.Context, err error) *subV1.GenericErrorStatusCode {
	return &subV1.GenericErrorStatusCode{}
}

package container

import (
	"net"
	"net/http"

	"github.com/goNiki/Subscription-service/internal/infrastructure/config"
	"github.com/goNiki/Subscription-service/internal/infrastructure/db"
	"github.com/goNiki/Subscription-service/internal/infrastructure/logger"
	"github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
)

type Container struct {
	Config              *config.Config
	Log                 logger.Logger
	SubscriptionsServer *subscriptions.Server
	Server              *http.Server
	DB                  *db.DB
}

func NewContainer(configpath string) (*Container, error) {

	c := &Container{}
	var err error

	c.Config, err = config.InitConfig(configpath)
	if err != nil {
		return &Container{}, err
	}

	c.Log = *logger.InitLogger(c.Config.Server.Env)

	c.SubscriptionsServer, err = subscriptions.NewServer(subscriptions.UnimplementedHandler{})
	if err != nil {
		return &Container{}, err
	}

	c.Server = &http.Server{
		Addr:        net.JoinHostPort(c.Config.Server.Host, c.Config.Server.Port),
		ReadTimeout: c.Config.Server.Timeout,
	}

	c.DB, err = db.NewDB(&c.Config.DB)
	if err != nil {
		return &Container{}, err
	}

	return c, nil
}

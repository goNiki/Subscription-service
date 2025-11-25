package container

import (
	"net"
	"net/http"

	v1 "github.com/goNiki/Subscription-service/internal/api/subscriptions/v1"
	"github.com/goNiki/Subscription-service/internal/infrastructure/config"
	"github.com/goNiki/Subscription-service/internal/infrastructure/db"
	"github.com/goNiki/Subscription-service/internal/infrastructure/logger"
	"github.com/goNiki/Subscription-service/internal/infrastructure/migrator"
	subscriptionsrepo "github.com/goNiki/Subscription-service/internal/repository/subscriptions"
	subscrtiptionsservice "github.com/goNiki/Subscription-service/internal/service/subscrtiptions"
	"github.com/goNiki/Subscription-service/shared/pkg/openapi/subscriptions/v1"
)

const MigDir = "./migrators"

type Container struct {
	Config               *config.Config
	Log                  logger.Logger
	DB                   *db.DB
	Migration            *migrator.Migrator
	SubscriptionsRepo    *subscriptionsrepo.Repository
	SubscriptionsService *subscrtiptionsservice.Service
	SubscriptionsApi     *v1.Api
	SubscriptionsServer  *subscriptions.Server
	Server               *http.Server
}

func NewContainer(configpath string) (*Container, error) {

	c := &Container{}
	var err error

	c.Config, err = config.InitConfig()
	if err != nil {
		return &Container{}, err
	}

	c.Log = *logger.InitLogger(c.Config.ServerConfig.Env)

	c.DB, err = db.NewDB(&c.Config.DBConfig)
	if err != nil {
		return &Container{}, err
	}

	c.Migration = migrator.NewMigrator(c.DB.Pool, MigDir)
	c.Migration.Up()

	c.SubscriptionsRepo = subscriptionsrepo.NewSubscriptionsRepo(c.DB.Pool)

	c.SubscriptionsService = subscrtiptionsservice.NewSubscriptionsService(c.SubscriptionsRepo)

	c.SubscriptionsApi = v1.NewSubscriptionsApi(c.Log.Log, c.SubscriptionsService)

	c.SubscriptionsServer, err = subscriptions.NewServer(c.SubscriptionsApi)
	if err != nil {
		return &Container{}, err
	}

	c.Server = &http.Server{
		Addr:        net.JoinHostPort(c.Config.ServerConfig.Host, c.Config.ServerConfig.Port),
		ReadTimeout: c.Config.ServerConfig.Timeout,
	}

	return c, nil
}

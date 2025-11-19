package db

import (
	"context"
	"fmt"
	"time"

	"github.com/goNiki/Subscription-service/internal/infrastructure/config"
	"github.com/goNiki/Subscription-service/models/errorapp"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

func NewDB(dbconfig *config.DBConfig) (*DB, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbconfig.User, dbconfig.Password, dbconfig.Host, dbconfig.Port, dbconfig.Name, dbconfig.Sslmode)

	PoolCfg, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return &DB{}, fmt.Errorf("%w: %v", errorapp.ErrParceConfigDB, err)
	}

	PoolCfg.MaxConns = dbconfig.MaxConns
	PoolCfg.MinConns = dbconfig.MinConns
	PoolCfg.MaxConnLifetime = dbconfig.MaxConnLifeTime
	PoolCfg.MaxConnIdleTime = dbconfig.MaxConnIdleTime
	PoolCfg.HealthCheckPeriod = dbconfig.HealthCheckPeriod

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, PoolCfg)
	if err != nil {
		return &DB{}, fmt.Errorf("%w: %v", errorapp.ErrCreateConnectionDB, err)
	}

	if err := pool.Ping(ctx); err != nil {
		return &DB{}, fmt.Errorf("%w: %v", errorapp.ErrPingDB, err)
	}

	return &DB{
		Pool: pool,
	}, nil
}

func (d *DB) CloseDB() {
	if d.Pool != nil {
		d.Pool.Close()
	}
}

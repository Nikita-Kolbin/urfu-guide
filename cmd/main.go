package main

import (
	"context"
	"fmt"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/api"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/cache/redis"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/config"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/repository"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/service"
	"github.com/Nikita-Kolbin/urfu-guide/internal/app/storage/minio"
	"github.com/Nikita-Kolbin/urfu-guide/internal/pkg/httpserver"
	"github.com/Nikita-Kolbin/urfu-guide/internal/pkg/logger"
)

// @title           URFU-Guide
// @version         1.0

// @host      158.160.182.170:8082
// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Token

func main() {
	ctx := context.Background()
	if err := initMain(ctx); err != nil {
		logger.Error(ctx, "run service failed", "err", err)
	}
}

func initMain(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return fmt.Errorf("init config failed: %w", err)
	}

	_ = cfg

	repo, err := repository.New(ctx, &cfg.Postgres)
	if err != nil {
		return fmt.Errorf("init reposytory failed: %w", err)
	}
	defer repo.Close(ctx)

	stg, err := minio.New(ctx, cfg.Minio.HostPort, cfg.Minio.Username, cfg.Minio.Password, cfg.Minio.UseSSL)
	if err != nil {
		return fmt.Errorf("init storage failed: %w", err)
	}

	cache, err := redis.NewClient(ctx, cfg.Redis.HostPort, cfg.Redis.Password)
	if err != nil {
		return fmt.Errorf("init cache failed: %w", err)
	}
	defer cache.Close()

	srv := service.New(repo, stg, cache, cfg.JWTSecret, cfg.ServerHostPort)

	r := api.New(ctx, srv, cfg.Listener.GetHostPort())

	server := httpserver.New(
		cfg.Listener.GetHostPort(), r,
		cfg.Listener.ReadTimeout,
		cfg.Listener.WriteTimeout,
		cfg.Listener.IdleTimeout,
	)

	logger.Info(ctx, "starting http server", "host_port", cfg.Listener.GetHostPort())
	if err = server.Run(); err != nil {
		return fmt.Errorf("failed run server: %w", err)
	}

	return nil
}

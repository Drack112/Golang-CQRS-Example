package server

import (
	"context"
	"errors"
	"net/http"

	echoserver "github.com/Drack112/Golang-GQRS-Example/internal/pkg/http/echo/server"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func RunServers(lc fx.Lifecycle, log logger.ILogger, e *echo.Echo, ctx context.Context, cfg *config.Config) error {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := echoserver.RunHttpServer(ctx, e, log, cfg.Echo); !errors.Is(err, http.ErrServerClosed) {
					log.Fatalf("error running http server: %v", err)
				}
			}()

			e.GET("/", func(c echo.Context) error {
				return c.String(http.StatusOK, config.GetMicroserviceName(cfg.ServiceName))
			})

			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Infof("all servers shutdown gracefully...")
			return nil
		},
	})

	return nil
}

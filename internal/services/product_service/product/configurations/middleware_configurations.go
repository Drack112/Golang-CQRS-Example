package configurations

import (
	"strings"

	echomiddleware "github.com/Drack112/Golang-GQRS-Example/internal/pkg/http/echo/middleware"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/otel"
	otelmiddleware "github.com/Drack112/Golang-GQRS-Example/internal/pkg/otel/middleware"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/constants"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func ConfigMiddlewares(e *echo.Echo, jaegerCfg *otel.JaegerConfig) {
	e.HideBanner = false
	e.Use(middleware.Logger())

	e.HTTPErrorHandler = middlewares.ProblemDetailsHandler
	e.Use(otelmiddleware.EchoTracerMiddleware(jaegerCfg.ServiceName))

	e.Use(echomiddleware.CorrelationsIdMiddleware)
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: constants.GzipLevel,
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Request().URL.Path, "swagger")
		},
	}))
	e.Use(middleware.BodyLimit(constants.BodyLimit))
}

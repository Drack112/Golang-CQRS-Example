package configurations

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	getting_products "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/gettings_products/v1/endpoints"
)

func ConfigEndPoints(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	getting_products.Maproute(validator, log, echo, ctx)
}

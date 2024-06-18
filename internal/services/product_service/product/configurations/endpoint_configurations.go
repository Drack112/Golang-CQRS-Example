package configurations

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	creating_product "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/creating_product/v1/endpoints"
	gettings_products "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/gettings_products/v1/endpoints"
)

func ConfigEndPoints(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	gettings_products.MapRoute(validator, log, echo, ctx)
	creating_product.MapRoute(validator, log, echo, ctx)
}

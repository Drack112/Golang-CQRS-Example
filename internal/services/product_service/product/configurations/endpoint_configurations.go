package configurations

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	creating_product "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/creating_product/v1/endpoints"
	deleting_product "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/deleting_product/v1/endpoints"
	getting_product_by_id "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/getting_product_by_id/v1/endpoints"
	gettings_products "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/gettings_products/v1/endpoints"
	searching_product "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/searching_product/v1/endpoints"
	updating_product "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/updating_product/v1/endpoints"
)

func ConfigEndPoints(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	gettings_products.MapRoute(validator, log, echo, ctx)
	creating_product.MapRoute(validator, log, echo, ctx)
	getting_product_by_id.MapRoute(validator, log, echo, ctx)
	searching_product.MapRoute(validator, log, echo, ctx)
	deleting_product.MapRoute(validator, log, echo, ctx)
	updating_product.MapRoute(validator, log, echo, ctx)
}

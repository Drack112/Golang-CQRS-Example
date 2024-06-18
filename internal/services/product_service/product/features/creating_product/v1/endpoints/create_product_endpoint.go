package endpoints

import (
	"context"
	"net/http"

	echomiddleware "github.com/Drack112/Golang-GQRS-Example/internal/pkg/http/echo/middleware"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
	"github.com/pkg/errors"

	commandsv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/creating_product/v1/commands"
	dtosv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/creating_product/v1/dtos"
)

func MapRoute(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
    group := echo.Group("/api/v1/products")
    group.POST("", createProduct(validator, log, ctx), echomiddleware.ValidateBearerToken())
}

func createProduct(validator *validator.Validate, log logger.ILogger, ctx context.Context) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := &dtosv1.CreateProductRequestDto{}
		if err := c.Bind(request); err != nil {
			badRequestErr := errors.Wrap(err, "[createProductEndpoint_handler.Bind] error in the binding request")
			log.Error(badRequestErr)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		command := commandsv1.NewCreateProduct(request.Name, request.Description, request.Price, request.InventoryId, request.Count)
		if err := validator.StructCtx(ctx, command); err != nil {
			validationErr := errors.Wrap(err, "[createProductEndpoint_handler.StructCtx] command validation failed")
			log.Error(validationErr)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		result, err := mediatr.Send[*commandsv1.CreateProduct, *dtosv1.CreateProductResponseDto](ctx, command)
		if err != nil {
			log.Errorf("(CreateProduct.Handle) id: {%s}, err: {%v}", command.ProductID, err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		log.Infof("(product created) id: {%s}", command.ProductID)
		return c.JSON(http.StatusCreated, result)
	}
}

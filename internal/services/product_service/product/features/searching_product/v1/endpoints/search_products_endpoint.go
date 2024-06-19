package endpoints

import (
	"context"
	"net/http"

	echomiddleware "github.com/Drack112/Golang-GQRS-Example/internal/pkg/http/echo/middleware"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"
	dtosv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/searching_product/v1/dtos"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/mehdihadeli/go-mediatr"
)

func MapRoute(validator *validator.Validate, log logger.ILogger, echo *echo.Echo, ctx context.Context) {
	group := echo.Group("/api/v1/products")
	group.GET("/search", searchProducts(validator, log, ctx), echomiddleware.ValidateBearerToken())
}

func searchProducts(validator *validator.Validate, log logger.ILogger, ctx context.Context) echo.HandlerFunc {

	return func(c echo.Context) error {
		listQuery, err := utils.GetListQueryFromCtx(c)

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		request := &dtosv1.SearchProductsRequestDto{ListQuery: listQuery}
		if err := c.Bind(request); err != nil {
			log.Warn("Bind", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		query := &dtosv1.SearchProductsRequestDto{SearchText: request.SearchText, ListQuery: request.ListQuery}
		if err := validator.StructCtx(ctx, query); err != nil {
			log.Errorf("(validate) err: {%v}", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		queryResult, err := mediatr.Send[*dtosv1.SearchProductsRequestDto, *dtosv1.SearchProductsResponseDto](ctx, query)
		if err != nil {
			log.Warn("SearchProducts", err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, queryResult)
	}

}

package dtos

import (
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/dtos"
)

type SearchProductsResponseDto struct {
	Products *utils.ListResult[*dtos.ProductDto]
}

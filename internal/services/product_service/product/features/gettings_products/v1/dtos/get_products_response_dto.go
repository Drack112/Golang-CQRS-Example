package dtos

import (
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/dtos"
)

type GetProductsResponseDto struct {
	Products *utils.ListResult[*dtos.ProductDto]
}

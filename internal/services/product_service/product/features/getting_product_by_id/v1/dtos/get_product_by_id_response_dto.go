package dtos

import "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/dtos"

type GetProductByIdResponseDto struct {
	Product *dtos.ProductDto `json:"product"`
}

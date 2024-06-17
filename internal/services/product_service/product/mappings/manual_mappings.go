package mappings

import (
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/dtos"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/models"
)

func ProductToProductResponseDto(product *models.Product) *dtos.ProductDto {
	return &dtos.ProductDto{
		ProductId:   product.ProductId,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

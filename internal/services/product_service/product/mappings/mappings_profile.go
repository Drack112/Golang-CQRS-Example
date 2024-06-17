package mappings

import (
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/mapper"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/dtos"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/models"
)

func ConfigureMappings() error {
	err := mapper.CreateMap[*models.Product, *dtos.ProductDto]()
	if err != nil {
		return err
	}

	return nil
}

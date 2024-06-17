package dtos

import "github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"

type GetProductsRequestDto struct {
	*utils.ListQuery
}

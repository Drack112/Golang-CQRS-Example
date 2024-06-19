package dtos

import "github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"

type SearchProductsRequestDto struct {
	SearchText       string `query:"search" json:"search"`
	*utils.ListQuery `json:"listQuery"`
}

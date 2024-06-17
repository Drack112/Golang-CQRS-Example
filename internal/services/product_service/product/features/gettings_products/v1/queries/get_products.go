package queries

import "github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"

type GetProducts struct {
	*utils.ListQuery
}

func NewGetProducts(query *utils.ListQuery) *GetProducts {
	return &GetProducts{ListQuery: query}
}

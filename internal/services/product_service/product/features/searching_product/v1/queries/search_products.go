package queries

import "github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"

type SearchProducts struct {
	SearchText string `validate:"required"`
	*utils.ListQuery
}

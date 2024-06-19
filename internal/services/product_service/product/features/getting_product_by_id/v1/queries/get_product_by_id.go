package queries

import uuid "github.com/satori/go.uuid"

type GetProductById struct {
	ProductID uuid.UUID `validate:"required"`
}

func NewGetProductByid(productId uuid.UUID) *GetProductById {
	return &GetProductById{ProductID: productId}
}

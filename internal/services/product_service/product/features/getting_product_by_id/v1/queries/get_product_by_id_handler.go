package queries

import (
	"context"
	"fmt"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/mapper"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/contracts"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/dtos"
	dtosv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/getting_product_by_id/v1/dtos"
	"github.com/pkg/errors"
)

type GetProductByIdHandler struct {
	log               logger.ILogger
	rabbitMqPublisher rabbitmq.IPublisher
	productRepository contracts.ProductRepository
	ctx               context.Context
	grpcClient        grpc.GrpcClient
}

func NewGetProductByIdHandler(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher, productRepository contracts.ProductRepository, ctx context.Context, grpcClient grpc.GrpcClient) *GetProductByIdHandler {
	return &GetProductByIdHandler{
		log:               log,
		rabbitMqPublisher: rabbitmqPublisher,
		productRepository: productRepository,
		ctx:               ctx,
		grpcClient:        grpcClient,
	}
}

func (q *GetProductByIdHandler) Handle(ctx context.Context, query *GetProductById) (*dtosv1.GetProductByIdResponseDto, error) {
	product, err := q.productRepository.GetProductById(ctx, query.ProductID)
	if err != nil {
		notFoundErr := errors.Wrap(err, fmt.Sprintf("product with id %s not found", query.ProductID))
		return nil, notFoundErr
	}

	productDtos, err := mapper.Map[*dtos.ProductDto](product)
	if err != nil {
		return nil, err
	}

	return &dtosv1.GetProductByIdResponseDto{Product: productDtos}, nil
}

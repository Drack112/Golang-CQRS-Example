package queries

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/contracts"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/dtos"
	dtosv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/gettings_products/v1/dtos"
)

type GetProductsHandler struct {
	log               logger.ILogger
	rabbitMqPublisher rabbitmq.IPublisher
	productRepository contracts.ProductRepository
	ctx               context.Context
	grpcClient        grpc.GrpcClient
}

func NewGetProductsHandler(log logger.ILogger, rabbitMqPublisher rabbitmq.IPublisher, productRepository contracts.ProductRepository, ctx context.Context, grpcClient grpc.GrpcClient) *GetProductsHandler {
	return &GetProductsHandler{
		log:               log,
		rabbitMqPublisher: rabbitMqPublisher,
		productRepository: productRepository,
		ctx:               ctx,
		grpcClient:        grpcClient,
	}
}

func (c *GetProductsHandler) Handle(ctx context.Context, query *GetProducts) (*dtosv1.GetProductsResponseDto, error) {
	products, err := c.productRepository.GetAllProducts(ctx, query.ListQuery)
	if err != nil {
		return nil, err
	}
	listResultDto, err := utils.ListResultTiListResultDto[*dtos.ProductDto](products)
	if err != nil {
		return nil, err
	}

	return &dtosv1.GetProductsResponseDto{Products: listResultDto}, nil
}

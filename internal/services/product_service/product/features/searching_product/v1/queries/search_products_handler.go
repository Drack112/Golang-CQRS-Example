package queries

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/utils"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/contracts"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/dtos"
	dtosv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/searching_product/v1/dtos"
)

type SearchProductsHandler struct {
	log               logger.ILogger
	rabbitMqPublisher rabbitmq.IPublisher
	productRepository contracts.ProductRepository
	ctx               context.Context
	grpcClient        grpc.GrpcClient
}

func NewSearchProductsHandler(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.ProductRepository, ctx context.Context, grpcClient grpc.GrpcClient) *SearchProductsHandler {
	return &SearchProductsHandler{log: log, productRepository: productRepository, ctx: ctx, rabbitMqPublisher: rabbitmqPublisher, grpcClient: grpcClient}
}

func (c *SearchProductsHandler) Handle(ctx context.Context, query *SearchProducts) (*dtosv1.SearchProductsResponseDto, error) {

	products, err := c.productRepository.SearchProducts(ctx, query.SearchText, query.ListQuery)
	if err != nil {
		return nil, err
	}

	listResultDto, err := utils.ListResultTiListResultDto[*dtos.ProductDto](products)
	if err != nil {
		return nil, err
	}

	return &dtosv1.SearchProductsResponseDto{Products: listResultDto}, nil
}

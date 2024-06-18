package commands

import (
	"context"
	"encoding/json"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/mapper"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/contracts"
	dtosv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/creating_product/v1/dtos"
	eventsv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/creating_product/v1/events"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/models"
)

type CreateProductHandler struct {
    log logger.ILogger
    rabbitMqPublisher rabbitmq.IPublisher
    productRepository contracts.ProductRepository
    ctx context.Context
    grpcClient grpc.GrpcClient
}
func NewCreateProductHandler(log logger.ILogger, rabbitMqPublisher rabbitmq.IPublisher, productRepository contracts.ProductRepository, ctx context.Context, grpcClient grpc.GrpcClient) *CreateProductHandler {
    return &CreateProductHandler{log: log, productRepository: productRepository, ctx: ctx, rabbitMqPublisher: rabbitMqPublisher, grpcClient: grpcClient}
}

func (c *CreateProductHandler) Handle(ctx context.Context, command *CreateProduct) (*dtosv1.CreateProductResponseDto, error) {

	product := &models.Product{
		ProductId:   command.ProductID,
		Name:        command.Name,
		Description: command.Description,
		Price:       command.Price,
		InventoryId: command.InventoryId,
		Count:       command.Count,
		CreatedAt:   command.CreatedAt,
	}

	createdProduct, err := c.productRepository.CreateProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	evt, err := mapper.Map[*eventsv1.ProductCreated](createdProduct)
	if err != nil {
		return nil, err
	}

	err = c.rabbitMqPublisher.PublishMessage(evt)
	if err != nil {
		return nil, err
	}

	response := &dtosv1.CreateProductResponseDto{ProductId: product.ProductId}
	bytes, _ := json.Marshal(response)

	c.log.Info("CreateProductResponseDto", string(bytes))

	return response, nil
}
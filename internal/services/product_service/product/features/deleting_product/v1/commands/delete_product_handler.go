package commands

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/contracts"
	eventsv1 "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/deleting_product/v1/events"
	"github.com/mehdihadeli/go-mediatr"
)

type DeleteProductHandler struct {
	log               logger.ILogger
	rabbitMqPublisher rabbitmq.IPublisher
	productRepository contracts.ProductRepository
	ctx               context.Context
	grpcClient        grpc.GrpcClient
}

func NewDeleteProductHandler(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.ProductRepository, ctx context.Context, grpcClient grpc.GrpcClient) *DeleteProductHandler {
	return &DeleteProductHandler{log: log, productRepository: productRepository, ctx: ctx, rabbitMqPublisher: rabbitmqPublisher, grpcClient: grpcClient}
}

func (c *DeleteProductHandler) Handle(ctx context.Context, command *DeleteProduct) (*mediatr.Unit, error) {
	if err := c.productRepository.DeleteProductByID(ctx, command.ProductID); err != nil {
		return nil, err
	}

	err := c.rabbitMqPublisher.PublishMessage(eventsv1.ProductDeleted{
		ProductId: command.ProductID,
	})
	if err != nil {
		return nil, err
	}

	c.log.Info("DeleteProduct successfully executed")
	return &mediatr.Unit{}, err
}

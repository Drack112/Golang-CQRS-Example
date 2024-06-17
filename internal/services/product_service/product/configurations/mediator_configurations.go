package configurations

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/contracts"
	gettingproductsv1dtos "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/gettings_products/v1/dtos"
	gettingproductsv1queries "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/gettings_products/v1/queries"
	"github.com/mehdihadeli/go-mediatr"
)

func ConfigProductsMediator(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.ProductRepository, ctx context.Context, grpcClient grpc.GrpcClient) error {

	//https://stackoverflow.com/questions/72034479/how-to-implement-generic-interfaces
	err := mediatr.RegisterRequestHandler[*gettingproductsv1queries.GetProducts, *gettingproductsv1dtos.GetProductsResponseDto](gettingproductsv1queries.NewGetProductsHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	return nil
}

package configurations

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/contracts"
	creatingproductv1commands "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/creating_product/v1/commands"
	creatingproductv1dtos "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/creating_product/v1/dtos"
	gettingproductbyidv1dtos "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/getting_product_by_id/v1/dtos"
	gettingproductbyidv1queries "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/getting_product_by_id/v1/queries"
	gettingproductsv1dtos "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/gettings_products/v1/dtos"
	gettingproductsv1queries "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/gettings_products/v1/queries"
	searchingproductsv1dtos "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/searching_product/v1/dtos"
	searchingproductsv1queries "github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/features/searching_product/v1/queries"
	"github.com/mehdihadeli/go-mediatr"
)

func ConfigProductsMediator(log logger.ILogger, rabbitmqPublisher rabbitmq.IPublisher,
	productRepository contracts.ProductRepository, ctx context.Context, grpcClient grpc.GrpcClient) error {

	//https://stackoverflow.com/questions/72034479/how-to-implement-generic-interfaces
	err := mediatr.RegisterRequestHandler[*creatingproductv1commands.CreateProduct, *creatingproductv1dtos.CreateProductResponseDto](creatingproductv1commands.NewCreateProductHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*gettingproductsv1queries.GetProducts, *gettingproductsv1dtos.GetProductsResponseDto](gettingproductsv1queries.NewGetProductsHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*gettingproductbyidv1queries.GetProductById, *gettingproductbyidv1dtos.GetProductByIdResponseDto](gettingproductbyidv1queries.NewGetProductByIdHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}

	err = mediatr.RegisterRequestHandler[*searchingproductsv1queries.SearchProducts, *searchingproductsv1dtos.SearchProductsResponseDto](searchingproductsv1queries.NewSearchProductsHandler(log, rabbitmqPublisher, productRepository, ctx, grpcClient))
	if err != nil {
		return err
	}
	return nil
}

package main

import (
	gormpsql "github.com/Drack112/Golang-GQRS-Example/internal/pkg/gorm_psql"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/http"
	echoserver "github.com/Drack112/Golang-GQRS-Example/internal/pkg/http/echo/server"
	httpclient "github.com/Drack112/Golang-GQRS-Example/internal/pkg/http_client"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/otel"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/config"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/configurations"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/repositories"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/mappings"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/models"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/server"
	"github.com/go-playground/validator/v10"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	fx.New(
		fx.Options(
			fx.Provide(
				config.InitConfig,
				logger.InitLogger,
				http.NewContext,
				echoserver.NewEchoServer,
				grpc.NewGrpcClient,
				gormpsql.NewGorm,
				otel.TracerProvider,
				httpclient.NewHttpClient,
				repositories.NewPostgresProductRepository,
				rabbitmq.NewRabbitMQConn,
				rabbitmq.NewPublisher,
				validator.New,
			),
			fx.Invoke(server.RunServers),
			fx.Invoke(configurations.ConfigMiddlewares),
			fx.Invoke(configurations.ConfigSwagger),
			fx.Invoke(func(gorm *gorm.DB) error {
				return gormpsql.Migrate(gorm, &models.Product{})
			}),
			fx.Invoke(mappings.ConfigureMappings),
			fx.Invoke(configurations.ConfigEndPoints),
			fx.Invoke(configurations.ConfigProductsMediator),
		),
	).Run()
}

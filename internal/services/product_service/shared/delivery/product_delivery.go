package delivery

import (
	"context"

	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/grpc"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/logger"
	"github.com/Drack112/Golang-GQRS-Example/internal/pkg/rabbitmq"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/config"
	"github.com/Drack112/Golang-GQRS-Example/internal/services/product_service/product/data/contracts"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type ProductDeliveryBase struct {
	Log               logger.ILogger
	Cfg               *config.Config
	RabbitmqPublisher rabbitmq.IPublisher
	ConnRabbitmq      *amqp.Connection
	HttpClient        *resty.Client
	JaegerTracer      trace.Tracer
	Gorm              *gorm.DB
	Echo              *echo.Echo
	GrpcClient        grpc.GrpcClient
	ProductRepository contracts.ProductRepository
	Ctx               context.Context
}

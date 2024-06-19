.PHONY:

### for running all commands we need bash command lien ###

## choco install make
# ==============================================================================
# Run Services
run_products_service:
	cd internal/services/product_service/ && 	go run ./cmd/main.go

## choco install protoc
## go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
## go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# ==============================================================================
# Proto Identity Service

## grpc-client
proto_identities_get_user_by_id_service:
	@echo Generating identity_service proto
	protoc --go_out=./internal/services/product_service/product/grpc_client/protos --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=./internal/services/product_service/product/grpc_client/protos ./internal/services/product_service/product/grpc_client/protos/*.proto

## go install github.com/swaggo/swag/cmd/swag@v1.8.3
# Swagger products Service  #https://github.com/swaggo/swag/issues/817
# ==============================================================================

swagger_products:
	@echo Starting swagger generating
	swag init -g ./internal/services/product_service/cmd/main.go -o ./internal/services/product_service/docs --exclude ./internal/services/identity_service, ./internal/services/inventory_service
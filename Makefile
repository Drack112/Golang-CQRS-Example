.PHONY:

### for running all commands we need bash command lien ###

## choco install make
# ==============================================================================
# Run Services
run_products_service:
	cd internal/services/product_service/ && 	go run ./cmd/main.go

## go install github.com/swaggo/swag/cmd/swag@v1.8.3
# Swagger products Service  #https://github.com/swaggo/swag/issues/817
# ==============================================================================
swagger_products:
	@echo Starting swagger generating
	swag init -g ./internal/services/product_service/cmd/main.go -o ./internal/services/product_service/docs --exclude ./internal/services/identity_service, ./internal/services/inventory_service

swagger_identities:
	@echo Starting swagger generating
	swag init -g ./internal/services/identity_service/cmd/main.go -o ./internal/services/identity_service/docs --exclude ./internal/services/product_service, ./internal/services/inventory_service

swagger_inventories:
	@echo Starting swagger generating
	swag init -g ./internal/services/inventory_service/cmd/main.go -o ./internal/services/inventory_service/docs --exclude ./internal/services/product_service, ./internal/services/identity_service
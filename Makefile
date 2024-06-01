BROKER_BINARY=broker-app

## broker: builds the broker binary as a linux executable
.Phony: broker
broker:
	@echo "Building broker binary..."
	cd ./broker-service && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Done."

## build: builds all services
.Phony: build
build: broker
	@echo "Stopping containers (if running) ..."
	docker compose down
	@echo "Building containers..."
	docker compose build
	@echo "Done."

## up: starts all the containers
.Phony: up
up:
	@echo "Starting containers..."
	docker compose up
	@echo "Done."

## down: shuts down all the containers
.Phony: down
down:
	@echo "Shutting down containers..."
	docker compose down
	@echo "Done."

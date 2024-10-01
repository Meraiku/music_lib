include .env

.ONESHELL:

build:
	@go build -o ./.bin/api ./cmd/api

run:build
	@echo "Swagger:		http://localhost:${PORT}/swagger/index.html"
	@./.bin/api

swagger:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@swag fmt
	@swag init -g cmd/api/main.go -o api/swagger

docker:build
	@docker compose up -d --build
	@echo "Swagger:		http://localhost:${PORT}/swagger/index.html"

stop:
	@docker compose down

test-unit:
	@go test $(shell go list ./... | grep -v /tests) -coverprofile=coverage.out
cover:
	go tool cover -html=coverage.out
test-integration:
	go test -race tests/integration/*.go

mocks:
	@go install go.uber.org/mock/mockgen@latest
	@mockgen -source=internal/repo/repo.go -destination=internal/repo/mocks/mock_repo.go


build_test:
	@go build -o ./.bin/info ./cmd/info

info:build_test
	@./.bin/info


up:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@cd ./sql/migrations;
	@goose postgres $(POSTGRES_DSN) up

down:
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@cd ./sql/migrations;
	@goose postgres $(POSTGRES_DSN) down
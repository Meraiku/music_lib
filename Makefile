

build:
	@go build -o ./.bin/api ./cmd/api

run:build
	@./.bin/api

swagger:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@swag init -g cmd/api/main.go -o api/swagger
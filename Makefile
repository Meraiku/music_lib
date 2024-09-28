

build:
	@go build -o ./.bin/api ./cmd/api

run:build
	@./.bin/api

swagger:
	@go install github.com/swaggo/swag/cmd/swag@latest
	@swag fmt
	@swag init -g cmd/api/main.go -o api/swagger

docker:build
	@sudo docker compose up --build


build_test:
	@go build -o ./.bin/info ./cmd/info

info:build_test
	@./.bin/info
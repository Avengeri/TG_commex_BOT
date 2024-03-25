export PWD ?= $(shell pwd)
export BIN_DIR ?= $(PWD)/.bin
export DEVOPS_DIR ?= $(PWD)/.devops

export COMPOSE_PROJECT_NAME ?= start
export DOCKER_COMPOSE = docker compose \
	-f ./docker-compose.yml \
	--env-file ./.env

ifeq ($(OS),Windows_NT)
	APP_EXT = .exe
endif

#################################################################### COMMON
download:
	go mod download
tidy:
	go mod tidy
upgrade:
	go get -u ./...
kill-port:
	npm kill-port 3000

#################################################################### DB
db:
	$(DOCKER_COMPOSE) up db -d
up:
	$(DOCKER_COMPOSE) up up
down:
	$(DOCKER_COMPOSE) up down
down1:
	$(DOCKER_COMPOSE) up down1
stop:
	$(DOCKER_COMPOSE) down

#################################################################### TRADING
app-build:
	go build -o ./cmd/start-go/api/${APP_EXT} cmd/start-go/api/main.go
app: app-build
	.bin/start-go${APP_EXT}
app-run:
	go run cmd/start-go/api/main.go

#################################################################### SWAGGER
swag-install:
	go install github.com/swaggo/swag/cmd/swag@latest
swag:
	swag init -g ./cmd/start-go/api/main.go

#################################################################### MIGRATIONS
migrate_up:
	docker run --rm -v ./migrations:/migrations --network host migrate/migrate -path=./migrations/ -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" up 1
migrate_down:
	docker run --rm -v ./migrations:/migrations --network host migrate/migrate -path=./migrations/ -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" down 1

#################################################################### START
start: db
	@sleep 1
	@make migrate_up
#################################################################### PROTOBUF

protoc:
	protoc -I=./proto --go_out=./proto/gen --go_opt=paths=source_relative \
    --go-grpc_out=./proto/gen --go-grpc_opt=paths=source_relative ./proto/cron.proto

#################################################################### BUILD
build_api_server:
	go build cmd/start-go/api/main_api_server.go
build_grpc_server:
	go build cmd/start-go/grpc/server/main_grpc_server.go
build_grpc_client:
	go build cmd/cron/grpc/client/main_grpc_client.go











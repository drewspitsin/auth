include .env

LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml
	
install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock@v3.1.3

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/gojuno/minimock/v3/cmd/minimock@v3.1.3

generate:
	make generate-user-api

generate-user-api:
	mkdir -p pkg/user_api_v1
	protoc --proto_path api/user_api \
	--go_out=pkg/user_api_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/user_api_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/user_api/user_api_v1.proto

build:
	GOOS=linux GOARCH=amd64 go build -o service_linux cmd/main.go

copy-to-server:
	scp service_linux root@89.104.117.12:
	scp .env root@89.104.117.12:

docker-build-and-push:
	docker buildx build --no-cache --platform linux/amd64 -t cr.selcloud.ru/ester-registry/test-server:v0.0.1 .
	docker login -u token -p CRgAAAAARPpNb1U5QUx83Ivvaz9cMQxXEdctTfKA cr.selcloud.ru/ester-registry
	docker push cr.selcloud.ru/ester-registry/test-server:v0.0.1

LOCAL_MIGRATION_DIR=$(MIGRATION_DIR)
LOCAL_MIGRATION_DSN="host=localhost port=$(PG_PORT) dbname=$(PG_DATABASE_NAME) user=$(PG_USER) password=$(PG_PASSWORD) sslmode=disable"

local-migration-status:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	$(LOCAL_BIN)/goose -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v




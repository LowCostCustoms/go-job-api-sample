ifndef $(GOPATH)
    GOPATH := $(shell go env GOPATH)
    export GOPATH
endif
export PATH := $(PATH):$(GOPATH)/bin

PROTO_TARGET_DIR := internal/api
SWAGGER_DIR := swagger
SOURCE_FILES := $(shell find . -type f -name "*.go")

.DEFAULT_GOAL := build

.PHONY: prerequisites
prerequisites:
	@echo "Downloading dependencies ..."
	@go get ./...
	@echo "Installing dependencies ..."
	@go install \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/golang/mock/mockgen \
		github.com/google/wire/cmd/wire

.PHONY: generate
generate: prerequisites
	@mkdir -p ${PROTO_TARGET_DIR} ${SWAGGER_DIR}
	@echo "Generating protocol stubs ..."
	@protoc \
		-Iproto \
		-Ithird_party/googleapis \
		--go_out ${PROTO_TARGET_DIR} \
		--go_opt paths=source_relative \
		--go-grpc_out ${PROTO_TARGET_DIR} \
		--go-grpc_opt paths=source_relative \
		--grpc-gateway_out ${PROTO_TARGET_DIR} \
		--grpc-gateway_opt paths=source_relative \
		--swagger_out ${SWAGGER_DIR} \
		./proto/*.proto
	@echo "Generating wire files ..."
	@wire gen ./...
	@echo "Generating mock files ..."
	@mockgen \
		-destination internal/test/mocks/http/http.go \
		-package http_mocks \
		net/http ResponseWriter,Handler
	@mockgen \
		-destination internal/test/mocks/repositories/repositories.go \
		-package repositories_mocks \
		-source internal/repositories/repositories.go JobRepository,JobScheduleRepository
	@mockgen \
		-destination internal/test/mocks/services/services.go \
		-package services_mocks \
		-source internal/services/services.go TransactionManager
	@mockgen \
		-destination internal/test/mocks/db/db.go \
		-package db_mocks \
		-source internal/db/db.go Transaction

.PHONY: build
build: generate $(SOURCE_FILES)
	@mkdir -p build
	@echo "Building application ..."
	@go build -o build/api .

.PHONY: migrate
migrate: build
	@echo "Running database migrations ..."
	@./build/api migrate

.PHONY: serve
serve: build
	@echo "Running web server ..."
	@./build/api serve

.PHONY: test
test: generate $(SOURCE_FILES)
	@echo "Running unit tests ..."
	@MIGRATIONS_LOCATION=file://$(shell pwd)/migrations/ \
		DATABASE_URL=postgres://postgres:postgres@localhost:5432/sched?sslmode=disable \
		go test -v -cover -coverprofile coverage.out ./...

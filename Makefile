include .env
export

.PHONY: help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: go-grpc-generate
go-grpc-generate: ## generate golang grpc code
	@if [ ! -d "filestorage/grpc" ]; then\
		mkdir -p filestorage/grpc;\
	fi
	@protoc file_storage.proto --go_out=./filestorage/grpc --go_opt=paths=source_relative
	@protoc file_storage.proto --go-grpc_out=./filestorage/grpc --go-grpc_opt=paths=source_relative

.PHONY: prisma-migrate-dev
prisma-migrate-dev: ## generate prisma migration locally
	@go run github.com/steebchen/prisma-client-go migrate dev

.PHONY: prisma-migrate-deploy
prisma-migrate-deploy: ## sync production database with migrations
	@go run github.com/steebchen/prisma-client-go migrate deploy

.PHONY: prisma-generate
prisma-generate: ## generate prisma go client
	@go run github.com/steebchen/prisma-client-go generate

.PHONY: test
test: ## run golang tests
	go test ./...

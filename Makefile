include .env
export

.PHONY: help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: go-grpc-generate
go-grpc-generate: ## generate golang grpc code
	@if [ ! -d "grpc" ]; then\
		mkdir -p grpc;\
	fi
	@protoc -I ./proto ./proto/*.proto --go_out=./grpc --go-grpc_out=./grpc

.PHONY: prisma-migrate-dev
prisma-migrate-dev:  DB_URL=postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DOCKER_PUBLISH_DB_PORT}/${DB_NAME}
prisma-migrate-dev: ## generate prisma migration locally
	@go run github.com/steebchen/prisma-client-go migrate dev

.PHONY: prisma-migrate-deploy
prisma-migrate-deploy: DB_URL=postgresql://${DB_USER}:${DB_PASSWORD}@localhost:${DOCKER_PUBLISH_DB_PORT}/${DB_NAME}
prisma-migrate-deploy: ## sync production database with migrations
	@go run github.com/steebchen/prisma-client-go migrate deploy

.PHONY: prisma-generate
prisma-generate: ## generate prisma go client
	@go run github.com/steebchen/prisma-client-go generate

.PHONY: test
test: ## run golang tests
	go test ./...

.PHONY: load-test-upload-file
load-test-upload-file: ## run load test upload file
	ghz --config ./load_test/upload_1MB.json localhost:${PORT}
	ghz --config ./load_test/upload_2MB.json localhost:${PORT}
	ghz --config ./load_test/upload_4MB.json localhost:${PORT}

.PHONY: docker-compose-prod-up
docker-compose-prod-up: ## start production docker compose
	@docker compose -f docker-compose.prod.yml up -d --build

.PHONY: docker-compose-prod-down
docker-compose-prod-down: ## stop production docker compose
	@docker compose -f docker-compose.prod.yml down -v

include ./deployments/env/.env

DEPLOY_BIN:=$(CURDIR)/deployments


#### MIGRATIONS

LOCAL_BIN:=$(CURDIR)/bin
MIGRATION_DSN="host=$(PG_HOST) port=$(PG_PORT) dbname=$(POSTGRES_DB) user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) sslmode=disable"

install-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0

migration-status:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} status -v

migration-up:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} up -v

migration-down:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} postgres ${MIGRATION_DSN} down -v

# коллекция баннеров
# make create-migration migration_name=banners_table

# коллекция слотов
# make create-migration migration_name=slots_table

# cоц-дем. групп пользователей
# make create-migration migration_name=social_group_table

# рабочая таблица баннеров в слотах
# make create-migration migration_name=banner_slot_table

# таблица статистики для баннеров в слотах
# make create-migration migration_name=banners_stats_table


create-migration:
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create 1_banners_table sql
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create 2_slots_table sql
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create 3_social_group_table sql
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create 4_banner_slot_table sql
	GOBIN=$(LOCAL_BIN) $(LOCAL_BIN)/goose -dir ${MIGRATION_DIR} create 5_banners_stats_table sql

#####


#### gRPC

grpc-install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.4.0

grpc-update-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

grpc-generate-api:
	mkdir -p pkg/rotation_v1
	protoc --proto_path api/rotation_v1 \
	--go_out=pkg/rotation_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/rotation_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/rotation_v1/rotation.proto

grpc-generate:
	grpc-generate-api
####

img-build:
	docker compose -f $(DEPLOY_BIN)/docker-compose.yaml -p $(COMPOSE_PROJECT_NAME) up -d --build

img-down:
	docker compose -f $(DEPLOY_BIN)/docker-compose.yaml -p $(COMPOSE_PROJECT_NAME) down -v

#############

build: img-build

run:

test:
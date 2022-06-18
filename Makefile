ifeq ($(ENVIRONMENT_NAME),docker)
	include .env.docker
else
	include .env.local
endif

.PHONY: create-migration
create-migration:
	echo Set NAME of migration file. eg: make create-migration NAME=create_user
	migrate create -ext sql -dir db/migration -seq $(NAME)

.PHONY: migrate-up
migrate-up:
	migrate -path db/migration -database "mysql://$(DB_SQL_URL)" -verbose up $(VERSION)

.PHONY: migrate-down-all
migrate-down-all:
	migrate -path db/migration -database "mysql://$(DB_SQL_URL)" -verbose down --all

.PHONY: migrate-down-version
migrate-down-version:
	migrate -path db/migration -database "mysql://$(DB_SQL_URL)" -verbose down $(VERSION)

.PHONY: migrate-drop-f
migrate-drop-f:
	migrate -path db/migration -database "mysql://$(DB_SQL_URL)" -verbose drop -f

.PHONY: seed-db
seed-db:
	go run cmd/seed/main.go

.PHONY: migrate-and-seed
migrate-and-seed: migrate-drop-f migrate-up seed-db

.PHONY: local
local:
	docker-compose --env-file ./.env.docker \
    	-f docker-compose.yml \
    	-f docker-compose.yml down

	docker-compose --env-file ./.env.docker \
	-f docker-compose.yml \
	-f docker-compose.yml build

	docker-compose --env-file ./.env.docker \
	-f docker-compose.yml \
	-f docker-compose.yml up
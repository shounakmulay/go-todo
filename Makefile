.PHONY: create-migration migrate-up migrate-down-all migrate-down-version migrate-drop-f seed-db migrate-and-seed

include .env.local
export

create-migration:
	echo Set NAME of migration file. eg: make create-migration NAME=create_user
	migrate create -ext sql -dir db/migration -seq $(NAME)

migrate-up:
	migrate -path db/migration -database "mysql://$(DB_SQL_URL)" -verbose up $(VERSION)

migrate-down-all:
	migrate -path db/migration -database "mysql://$(DB_SQL_URL)" -verbose down --all

migrate-down-version:
	migrate -path db/migration -database "mysql://$(DB_SQL_URL)" -verbose down $(VERSION)

migrate-drop-f:
	migrate -path db/migration -database "mysql://$(DB_SQL_URL)" -verbose drop -f

seed-db:
	go run cmd/seed/main.go

migrate-and-seed: migrate-drop-f migrate-up seed-db

local:
	go run cmd/server/main.go
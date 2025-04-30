include .env

DB_CONN=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_HOST_PORT}/${POSTGRES_DB}?sslmode=disable

.PHONY: create-migration migrate-up migrate-down

create-migration:
	@read -p "Enter the sequence name: " SEQ; \
		docker run --rm -v ./sql/migrations:/migrations migrate/migrate \
			create -ext sql -dir /migrations -seq $${SEQ}

migrate-up:
	@docker run --rm -v ./sql/migrations:/migrations --network host migrate/migrate \
		-path=/migrations -database "${DB_CONN}" up

migrate-down:
	@read -p "Number of migrations you want to rollback (default: 1): " NUM; NUM=$${NUM:-1}; \
	docker run --rm -it -v ./sql/migrations:/migrations --network host migrate/migrate \
		-path=/migrations -database "${DB_CONN}" down $${NUM}

build:
	@go build -o tmp/run .

run: build
	@./tmp/run

dev:
	@air -c .air.toml

fmt:
	@gofmt -l -s -w .

CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go
proto-gen:
	./scripts/gen-proto.sh

run:
	go run cmd/main.go

migrate-up:
	migrate -path migrations -database "postgresql://postgres:123@localhost:5432/product?sslmode=disable" -verbose up

migrate-down:
	migrate -path migrations -database "postgres://postgres:123@localhost:5432/product?sslmode=disable" -verbose down

migrate_file:
	migrate create -ext sql -dir migrations/ -seq product

migrate-dirty:
	migrate -path ./migrations/ -database "postgresql://postgres:123@localhost:5432/product?sslmode=disable" force 1

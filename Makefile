APP := treasury-api

.PHONY: run worker test lint tidy build proto

run:
	go run ./cmd/api

worker:
	go run ./cmd/worker

build:
	CGO_ENABLED=0 go build -o bin/$(APP)-api ./cmd/api
	CGO_ENABLED=0 go build -o bin/$(APP)-worker ./cmd/worker

lint:
	golangci-lint run ./...

test:
	go test ./...

proto:
	buf generate

tidy:
	go mod tidy

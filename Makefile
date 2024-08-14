.PHONY: build test build-docker run test-one

build:
	go build -o ./cmd/run.go

format:
	go fmt ./...

run:
	go run cmd/run.go

test:
	go test -v ./stlocks

bench:
	go test ./stlocks -bench .
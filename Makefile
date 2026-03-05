# Standard Makefile targets for building, testing, and maintaining the go-zod-port library

.PHONY: build test lint fmt vet bench coverage

build:
	go build -o go-zod-port ./...

echo "Build successful"

test:
	go test -race ./...

echo "Tests passed"

lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

bench:
	go test -bench=. -benchmem ./...

echo "Benchmarking complete"

coverage:
	go test -coverprofile=coverage.out -covermode=atomic ./...

echo "Coverage report generated"

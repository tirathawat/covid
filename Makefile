install:
	go mod tidy

dev:
	go run cmd/api/main.go

test: unit-test vet

unit-test:
	go test -tags=unit -race ./...

vet: 
	go vet ./...
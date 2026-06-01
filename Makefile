.PHONY: run test coverage

run:
	go run ./cmd/server

test:
	go test ./... -race

coverage:
	go test ./... -race -coverprofile=coverage.out
	go tool cover -html=coverage.out
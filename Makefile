build:
	@go build -o bin/Hocus cmd/main.go

run: build
	@./bin/Hocus

test:
	@go test -v ./...
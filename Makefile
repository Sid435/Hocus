build:
	go build -o bin/Hocus cmd/main.go

run:
	./bin/Hocus

test:
	go test -v ./...
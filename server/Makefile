.PHONY: build run test

build:
	go build -o webserver .

run: build
	./webserver -root .. -addr 127.0.0.1:8080

test:
	go vet ./... && go test ./...

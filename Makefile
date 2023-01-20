cmd=httpserver
app=

fmt:
	go fmt ./...

lint:fmt
	golint ./...

vet:
	go vet ./...

build:vet
	go build -o $(cmd) cmd/$(cmd)/main.go

run:build
	./httpserver

test:
	go test -v  ./...

.PHONY: fmt lint vet build run test
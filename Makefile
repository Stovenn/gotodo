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
	./$(app)

.PHONY: fmt lint vet build
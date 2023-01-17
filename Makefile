fmt:
	go fmt ./...

lint:fmt
	golint ./...

vet:
	go vet ./...

build:vet
	go build -o gotodo cmd/main.go

run:build
	./gotodo

.PHONY: fmt lint vet build
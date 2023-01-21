fmt:
	go fmt ./...

lint:fmt
	golint ./...

vet:
	go vet ./...

build:vet
	go build -o gotodo cmd/gotodo/main.go

run:
	./gotodo

test:
	go test -v  ./...

mock_repo:
	mockgen -package mockdb -destination internal/repositories/mock/store.go github.com/stovenn/gotodo/internal/core/ports TodoRepository

mock_service:
	mockgen -package mockservice -destination internal/core/services/todoservice/mock/service.go github.com/stovenn/gotodo/internal/core/ports TodoService

.PHONY: fmt lint vet build run test mock_repo mock_service
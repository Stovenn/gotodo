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

postgres:
	docker run --name psql -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -dp 5432:5432 postgres:13

createdb:
	docker exec -it psql createdb --username=postgres --owner=postgres gotodo

dropdb:
	docker exec -it psql dropdb --username=postgres gotodo

mig_name =

create_migration:
	 migrate create -ext sql -dir internal/repositories/psqlrepo/migrations -seq $(mig_name)

migrateup:
	migrate -path internal/repositories/psqlrepo/migrations -database "postgresql://postgres:password@localhost:5432/gotodo?sslmode=disable" -verbose up

migrate1up:
	migrate -path internal/repositories/psqlrepo/migrations -database "postgresql://postgres:password@localhost:5432/gotodo?sslmode=disable" -verbose up 1

migratedown:
	migrate -path internal/repositories/psqlrepo/migrations -database "postgresql://postgres:password@localhost:5432/gotodo?sslmode=disable" -verbose down

migrate1down:
	migrate -path internal/repositories/psqlrepo/migrations -database "postgresql://postgres:password@localhost:5432/gotodo?sslmode=disable" -verbose down 1


.PHONY: fmt lint vet build run test mock_repo mock_service postgres createdb dropdb create_migration migrateup migrate1up migratedown migrate1down
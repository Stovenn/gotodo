# GoTodo

## Installation
Clone the repository
```
git clone https://github.com/Stovenn/gotodo.git
```
Build binary (only unix system for the moment)
```
make build
```
Launch server
```
make run
```
## Usage
The server listens on port 8080 of host machine (can be configured in .env file)
### Create a todo
POST  http://localhost:8080/api/todos/
```
{
  "title": "new todo"
}
```
### List all todos
GET  http://localhost:8080/api/todos/
### Get a single todo
GET  http://localhost:8080/api/todos/{id}
### Update a todo
PUT  http://localhost:8080/api/todos/{id}
```
{
  "title": "updated todo",
  "completed": true,
  "order": 2,
}
```
### Partially update a todo (not implemented)
PATCH  http://localhost:8080/api/todos/{id}
```
{
  "title":"partially updated todo",
  "order": 2,
}
```
### Delete a todo
DELETE  http://localhost:8080/api/todos/{id}

## Tests
```
make test
```

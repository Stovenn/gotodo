# Build stage
FROM golang:1.19.5-alpine3.17 AS build
WORKDIR /app
COPY . .
RUN go build -o gotodo cmd/gotodo/main.go

# Run Stage
FROM alpine:3.17
WORKDIR /app
COPY --from=build /app/gotodo .
COPY app.env .

EXPOSE 8080
CMD ["/app/gotodo"]
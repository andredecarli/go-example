BINARY_NAME=app

.PHONY: build run up clean test go-run docker-compose-up

build:
	docker build -t $(BINARY_NAME):latest .

run:
	docker run --rm -p 8080:8080 $(BINARY_NAME):latest

up: build run

clean:
	docker rmi $(BINARY_NAME):latest

test:
	go test ./... -cover

go-run:
	go run ./cmd/main.go

docker-compose-up:
	docker-compose up --build

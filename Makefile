BINARY_NAME=app

.PHONY: build run up clean

build:
	docker build -t $(BINARY_NAME):latest .

run:
	docker run --rm -it $(BINARY_NAME):latest

up: build run

clean:
	docker rmi $(BINARY_NAME):latest

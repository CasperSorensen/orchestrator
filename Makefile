# Variables
BINARY_NAME = "orchestrator"
SRC_DIR=./cmd/main.go

all:build

build:
	go build -o $(BINARY_NAME) $(SRC_DIR)

run: clean build
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

tidy:
	go mod tidy
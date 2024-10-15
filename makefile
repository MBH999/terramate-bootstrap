# Define the binary name and output directory
BINARY_NAME=terramate-bootstrap
BIN_DIR=./bin
SRC=./main.go

# Create the bin directory if it doesn't exist
.PHONY: build clean

build: 
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC)

install:
	sudo go build -o /usr/local/bin/$(BINARY_NAME) $(SRC)
	

clean:
	rm -rf $(BIN_DIR)/*

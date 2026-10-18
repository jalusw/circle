BIN      := $(CURDIR)/bin
GOFILES  := $(shell find . -name '*.go')
DATABASE ?= circle.db

.PHONY: build test vet lint run-server run-cli clean

build:
	mkdir -p $(BIN)
	go build -o $(BIN)/circle ./cmd/circle
	go build -o $(BIN)/circle-server ./cmd/circle-server
	go build -o $(BIN)/circle-tui ./cmd/circle-tui

vet:
	go vet ./...

lint:
	golangci-lint run

scratch-db:
	rm -f $(DATABASE)

run-server:
	go run ./cmd/circle-server

clean:
	rm -rf $(BIN) $(DATABASE)

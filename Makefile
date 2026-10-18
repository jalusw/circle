BIN      := $(CURDIR)/bin
GOFILES  := $(shell find . -name '*.go')
DATABASE ?= cycle.db

.PHONY: build test vet lint run-server run-cli clean

build:
	mkdir -p $(BIN)
	go build -o $(BIN)/cycle ./cmd/cycle
	go build -o $(BIN)/cycle-server ./cmd/cycle-server
	go build -o $(BIN)/cycle-tui ./cmd/cycle-tui

vet:
	go vet ./...

lint:
	golangci-lint run

scratch-db:
	rm -f $(DATABASE)

run-server:
	go run ./cmd/cycle-server

clean:
	rm -rf $(BIN) $(DATABASE)

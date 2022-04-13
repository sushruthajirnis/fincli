SHELL := /bin/bash

# Default Go linker flags.
GO_LDFLAGS ?= -ldflags="-s -w"

FINCLI := ./bin/fincli

$(FINCLI):
	      CGO_ENABLED=0 GOOS=darwin go build -trimpath $(GO_LDFLAGS)  -o $@ . # POC apple

lint:
	golangci-lint run --fast

clean:
	rm -rf ./bin

test:
	go test

.PHONY: all
all: clean $(FINCLI)

.PHONY: install
install: all
		 zip -j bin/fincli.zip $(FINCLI) #homebrew find better


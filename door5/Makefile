.DEFAULT_GOAL := all

.PHONY: fmt vet build

all: clean build

fmt: 
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	mkdir -p bin
	go build -o bin/door5 ./door5.go

clean:
	go clean
	rm -drf bin


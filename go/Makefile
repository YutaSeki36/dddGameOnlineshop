# Makefile

all: download test build

build:
	go build -o app

test:
	go test -v ./...

download:
	go mod download

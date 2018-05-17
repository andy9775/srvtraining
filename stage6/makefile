SHELL := /bin/bash

all: crud

crud:
	cd "$$GOPATH/src/github.com/ardanlabs/srvtraining/stage6"
	docker build -t srvt-crud-amd64 -f dockerfile.crud .
	docker system prune -f

up:
	docker-compose up

down:
	docker-compose down

test:  
	source "$$GOPATH/src/github.com/ardanlabs/srvtraining/stage6"
	go test ./...
#!make
init:
	go install github.com/cosmtrek/air@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install gotest.tools/gotestsum@latest
	brew install golangci-lint

deps:
	go get
	go mod tidy
	go mod vendor

build:
	go build -v ./...

test:
	gotestsum --format testname

lint:
	golangci-lint run

gosec:
	gosec ./...

govulncheck:
	govulncheck ./...

security: gosec govulncheck

clean:
	go clean
	rm -f bin/*

check: build test lint security clean

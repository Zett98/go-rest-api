.Phony: build
build: 
	go build -v ./cmd/apiserver

.Phony: run
run:
	go run -race -v ./cmd/apiserver/main.go

.Phony: test
test: 
	go test -v -race -timeout 30s ./...	

.DEFAULT_GOAL := build	
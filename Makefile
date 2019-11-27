.Phony: build
build: 
	go build -v ./cmd/apiserver

.Phony: run
run:
	go build -v ./cmd/apiservser; ./apiserver

.Phony: test
test: 
	go test -v -race -timeout 30s ./...	

.DEFAULT_GOAL := build	
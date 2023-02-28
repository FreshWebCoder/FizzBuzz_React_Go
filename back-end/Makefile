.PHONY: all
all:
	make dependencies && \
	make lint && \
	make format

.PHONY: lint 
lint: 
	golangci-lint run --config=./build/golangci-lint/.golangci.yml

.PHONY: test 
test:
	sh build/testing/test.sh	

.PHONY: format 
format: 
	go fmt ./...

.PHONY: dependencies
dependencies:
	brew upgrade golangci-lint

.PHONY: prepare 
prepare:
	brew update && brew upgrade && brew install golangci-lint

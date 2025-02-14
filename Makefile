GO_FILES := $(wildcard *.go)

build:
	go build -o myBlockchainApp $(GO_FILES)

run: build
	./myBlockchainApp

fmt:
	go fmt ./...

test:
	go test -v ./...

clean:
	rm -f myBlockchainApp

.PHONY: build run fmt test clean
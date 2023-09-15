.PHONY: all stat

all: gofmt test

test:
	@echo "run test..."
	go clean -testcache
	go test ./... -cover

gofmt:
	@if [ -x "$$(command -v gofmt)" ]; then echo "run gofmt..." ; gofmt -s -w . ; else echo "gofmt not found (skipped)" ; fi

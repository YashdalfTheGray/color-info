.PHONY: all clean test run
all: build test

# go list is the canonical utility to find go files
GOFILES := $(shell go list -f '{{ join .GoFiles "\n" }}' ./...)

build: .bin-stamp
	go build -o bin/color-info
	chmod +x bin/color-info

test:
	go test -covermode=atomic -coverpkg=all ./...

run-bin: build
	./bin/color-info

run:
	go run .

# directories do werid things in make, so we can use a stamp
.bin-stamp:
	mkdir -p bin
	touch .bin-stamp

clean:
	rm -rf bin
	rm .bin-stamp

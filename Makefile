PACKAGES=$(shell go list ./... | grep -v 'vendor/' | grep -v 'template/')
BUILD_FLAGS = -ldflags "-X github.com/svaishnavy/cosmos-zone/version.GitCommit=`git rev-parse --short HEAD`"

all: get_tools get_vendor_deps pack pack_clean test install

get_tools:
	go get github.com/golang/dep/cmd/dep
	go get -u github.com/gobuffalo/packr/...

pack:
	packr build $(BUILD_FLAGS) -o bin/cosmos-zone ./main.go

build:
	go build $(BUILD_FLAGS) -o bin/cosmos-zone ./main.go

pack_clean:
	packr clean

get_vendor_deps:
	@rm -rf vendor/
	@dep ensure

test:
	@go test $(PACKAGES)

benchmark:
	@go test -bench=. $(PACKAGES)

install:
	go install github.com/svaishnavy/cosmos-zone

.PHONY: all build test benchmark


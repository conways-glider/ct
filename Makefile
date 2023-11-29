OS ?=
CGO_ENABLED ?=
GOOS ?=
GOARCH ?=
VERSION ?= 0.0.0

build:
	go clean
	mkdir -p bin
	go build -o ./bin/

release:
	go clean
	mkdir -p bin
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-s -w -X main.Version=${VERSION}" -o ./bin/

clean:
	go clean
	rm -rf bin
	rm -rf dist

strip:
	strip ./bin/ct

generate-test-files:
	./scripts/generate_test_files.sh

test:
	go test -v ./...

race-condition:
	go build -v -race ./...

cover:
	go test -covermode=count -coverprofile=count.out ./...

mod-tidy-check:
	go mod tidy
	git diff --exit-code

lint: lint-go shellcheck

lint-go:
	go clean
	golangci-lint run ./...

shellcheck:
	./scripts/shellcheck.sh

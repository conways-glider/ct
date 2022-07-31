OS ?=
CGO_ENABLED ?=
GOOS ?=
GOARCH ?=
RELEASE_VERSION ?= 0.0.0

build:
	go clean
	mkdir -p bin
	go build -o ./bin/ct

build-release:
	go clean
	mkdir -p bin
	mkdir -p dist
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -ldflags "-s -w -X main.Version=${RELEASE_VERSION}" -o ./bin/ct

strip:
	strip ./bin/ct

generate-test-files:
	./scripts/generate_test_files.sh

test:
	go test ./... -v

mod-tidy-check:
	go mod tidy
	git diff --exit-code

lint:
	go clean
	golangci-lint run ./...

shellcheck:
	./scripts/shellcheck.sh

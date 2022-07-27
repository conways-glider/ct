CGO_ENABLED ?=
GOOS ?=
GOARCH ?=
VERSION ?= 0.0.0

build:
	go clean
	mkdir -p bin
	go build -o ./bin/ct

build-release:
	go clean
	mkdir -p bin
	mkdir -p build
	CGO_ENABLED=${CGO_ENABLED} GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./bin/ct
	zip -r -j build/ct_${VERSION}_${GOOS}_${GOARCH}.zip ./bin

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

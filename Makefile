all: lint build test

build:
	go build ./...

install:
	./scripts/make-install.sh

lint: fmt vet staticcheck misspell

fmt:
	./scripts/gofmt.sh

vet:
	go vet ./check ./cmd/... ./download ./handlers ./tools/...
	go vet ./main.go

staticcheck:
	@[ -x "$(shell which staticcheck)" ] || go install honnef.co/go/tools/cmd/staticcheck
	staticcheck ./...

test:
	 go test -cover ./...

start:
	 go run main.go


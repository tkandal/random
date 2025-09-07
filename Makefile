GO ?= go
GOFMT ?= gofmt
GOLINT ?= golint
GOSEC ?= gosec
GOVULNCHECK ?= govulncheck
GOCRITIC ?= gocritic
STATICCHECK ?= staticcheck

SOURCES = $(wildcard *.go)

fmt:
	$(GOFMT) -w $(SOURCES)

lint:
	$(GOLINT) ./...

vet:
	$(GO) vet ./...

gosec:
	$(GOSEC) ./...

critic:
	$(GOCRITIC) check ./...

static:
	$(STATICCHECK) -tests ./...

test:
	$(GO) test -coverprofile=coverage.out ./...
	$(GO) test -json > test.json ./...

check:
	$(GOVULNCHECK) ./...

.PHONY:	fmt lint vet gosec critic static test check

## Install library for production
deps:
	go get -u ./...
.PHONY: deps

## Install library for development
devel-deps: deps
	GO111MODULE=off go get \
		golang.org/x/lint/golint \
		honnef.co/go/tools/staticcheck \
		github.com/kisielk/errcheck \
		golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow \
		github.com/securego/gosec/cmd/gosec \
		github.com/motemen/gobump/cmd/gobump \
		github.com/Songmu/make2help/cmd/make2help
.PHONY: devel-deps

## Execute unit test
test: deps
	go test -v -count=1 -cover ./...
.PHONY: test

## Output coverage of testing
cov:
	go test -count 1 -coverprofile=cover.out ./...
	go tool cover -html=cover.out
.PHONY: cov

## Lint
lint: devel-deps
	go vet ./...
	staticcheck ./...
	errcheck ./...
	gosec -quiet ./...
	golint -set_exit_status ./...
	shadow ./...
.PHONY: lint

## Help
help:
	@make2help --all
.PHONY: help

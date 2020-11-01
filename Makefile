GO=go

default: help

.PHONY: format
format: #> Format the codebase
	$(GO) fmt ./...
	goimports -w .

.PHONY: test
test: #> Run unit tests
	$(GO) test -v -race ./...

.PHONY: test-cover
test-cover: #> Run unit tests with coverage
	$(GO) test -v -race -coverprofile=coverage.out ./...
	$(GO) tool cover -html=coverage.out
	rm coverage.out

.PHONY: help
help:
	@ echo "Usable Commands:"
	@ grep -h -E '^[a-zA-Z0-9_-]+:.*?#> .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?#> "}; {printf "\033[36m%-17s:\033[0m %s\n", $$1, $$2}'

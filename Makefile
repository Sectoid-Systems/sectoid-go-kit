# Define variables
VERSION ?= v0.1.0



.PHONY: test
test: ## Run all tests on local
test:
	@go test ./...

.PHONY: clean
clean: ## Clean the binary
clean:
	@go clean

.PHONY: lint
lint: ## Run linter
lint:
	@golangci-lint run

.PHONY: fmt
fmt: ## Format target
fmt:
	@gofmt -w .

.PHONY: publish
publish: ## Publish target
publish: clean
publish: test
publish: fmt
publish: lint
publish:
	@if [ -z "$(VERSION)" ]; then echo "Error: VERSION is not set"; exit 1; fi
	@git tag $(VERSION)
	@GOPROXY=proxy.golang.org go list -m github.com/Sectoid-Systems/sectoid-go-kit@$(VERSION)

.PHONY: help
help: ## Help target
help:
	@printf "\n$$(tput bold)Available targets:$$(tput sgr0)\n"
	@grep -h -E '^[%$$\(\)$a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed -E 's|-common-default||g' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[92m%-20s\033[0m %s\n", $$1, $$2}'
	@printf "\n"

.DEFAULT_GOAL := help
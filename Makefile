.PHONY: test
test: ## Run all tests on local
test:
	@go test ./...

.PHONY: help
help: ## Help target
help:
	@printf "\n$$(tput bold)Available targets:$$(tput sgr0)\n"
	@grep -h -E '^[%$$\(\)$a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sed -E 's|-common-default||g' | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[92m%-20s\033[0m %s\n", $$1, $$2}'
	@printf "\n"

.DEFAULT_GOAL := help
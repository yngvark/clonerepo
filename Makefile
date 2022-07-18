.PHONY: help

help: ## Print this menu
	@grep -E '^[a-zA-Z_0-9-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


test: install-richgo ## -
	richgo test \
		-parallel 8 \
		-timeout 10m \
		./...

build: ## -
	mkdir -p ./build
	go build -o ./build/go-cmd-template

run: ## -
	go run ./...

fmt: install-gofumpt ## -
	go fmt ./...
	gofumpt -l -w .

lint: install-lint ## -
	golangci-lint run ./...

check: fmt lint ## Run fmt and lint (so you have one command use in your dev flow instead of two)

#
# Installations
#

# https://github.com/mvdan/gofumpt
install-gofumpt:
	@go install mvdan.cc/gofumpt@v0.3.1

# # https://golangci-lint.run/usage/install/#local-installation
install-lint:
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/9d9855c149a3d46410f0bf818ead38c9f445bbf1/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.46.2


# https://github.com/kyoh86/richgo
install-richgo:
	@go install github.com/kyoh86/richgo@v0.3.10

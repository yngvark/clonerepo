.PHONY: help
SHELL = bash
SOURCES := $(shell find ./cmd ./pkg -name "*.go")

help: ## Print this menu
	@grep -E '^[a-zA-Z_0-9-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

test: install-richgo ## -
	richgo test \
		-parallel 8 \
		-timeout 10m \
		./...

build: $(SOURCES) ## -
	mkdir -p ./build
	go build -o ./build/clonerepo

run: ## -
	go run ./...

fmt: install-gofumpt ## -
	go fmt ./...
	gofumpt -l -w .

lint: install-lint ## -
	golangci-lint run ./...

check: fmt lint test ## Run fmt and lint (so you have one command use in your dev flow instead of two)

#
# Installations
#

# https://github.com/mvdan/gofumpt
install-gofumpt:
ifneq ($(shell gofumpt --version), v0.3.1)
	@go install mvdan.cc/gofumpt@v0.3.1
endif


# # https://golangci-lint.run/usage/install/#local-installation
install-lint:
ifneq ($(shell golangci-lint version --format short), 1.46.2)
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/6f7f8aebbee04a8f6acff8bb37ae86746b9e5e0d/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.50.1
endif

# https://github.com/kyoh86/richgo
install-richgo:
	@go install github.com/kyoh86/richgo@v0.3.10

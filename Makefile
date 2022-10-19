GENERATED_DIR ?= openapiclient
SOURCE_FILES ?= $(shell find . -not -path "./$(GENERATED_DIR)/*" -type f -name '*.go' -print)

GOLANGCI_LINT_VERSION ?= 1.50.0

GIT_REVISION ?= $(shell git rev-parse --short HEAD)
GIT_TAG ?= $(shell git describe --tags --abbrev=0 | sed -e s/v//g)

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.DEFAULT_GOAL := help

.PHONY: generate
generate: ## run OpenAPI Generator
	rm -rf $(GENERATED_DIR)
	npx @openapitools/openapi-generator-cli generate \
		--input-spec spec.yaml \
		--generator-name go \
		--output $(GENERATED_DIR) \
		--package-name $(GENERATED_DIR) \
		--git-host github.com \
		--git-user-id ks6088ts \
		--git-repo-id soracom-sdk-go \
		--http-user-agent ks6088ts/soracom-sdk-go/$(GIT_TAG)
	cd $(GENERATED_DIR) && rm -rf .openapi-generator api .gitignore .openapi-generator-ignore .travis.yml git_push.sh go.mod go.sum

.PHONY: generate-diff-check
generate-diff-check: ## check if generated codes are the same as tracked ones
	git diff --exit-code --relative $(GENERATED_DIR)

.PHONY: install-deps-dev
install-deps-dev: ## install dependencies for development
	# https://golangci-lint.run/usage/install/#linux-and-windows
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v$(GOLANGCI_LINT_VERSION)

.PHONY: format
format: ## format codes
	gofmt -s -w $(SOURCE_FILES)

.PHONY: lint
lint: ## lint
	golangci-lint run -v

.PHONY: test
test: ## run tests
	go test -cover -v ./... -count=1

.PHONY: ci-test
ci-test: generate generate-diff-check install-deps-dev lint test ## ci test

.PHONY: test vet fmt lint new-problem help
PKG ?= ./...

test: ## Run tests (override with PKG=./problems/0001-two-sum/...)
	go test $(PKG)

vet: ## Run go vet
	go vet ./...

fmt: ## Format code
	gofmt -l -w .

lint: ## Run golangci-lint if installed
	golangci-lint run ./...

new-problem: ## Scaffold a problem: make new-problem NUM=0002 SLUG=add-two-numbers PKG=addtwonumbers FUNC=AddTwoNumbers
ifndef NUM
	$(error NUM is required, e.g. NUM=0002)
endif
ifndef SLUG
	$(error SLUG is required, e.g. SLUG=add-two-numbers)
endif
ifndef PKG
	$(error PKG is required, e.g. PKG=addtwonumbers)
endif
ifndef FUNC
	$(error FUNC is required, e.g. FUNC=AddTwoNumbers)
endif
	./scripts/new-problem.sh $(NUM) $(SLUG) $(PKG) $(FUNC)

help:
	@grep -E '^[a-z-]+:.*?##' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-12s %s\n", $$1, $$2}'

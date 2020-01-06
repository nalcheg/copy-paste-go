all: prepare run-tests purge

prepare: ## Prepare.
	@docker-compose up -d
	@docker-compose ps --all

run-app: ## Run app.
	go run ./main.go ./common-ops.go

run-tests: ## Run tests.
	go test -count=1 -v ./...

purge: ## Purge.
	@docker-compose rm -sf && docker-compose down
	@docker-compose ps --all

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

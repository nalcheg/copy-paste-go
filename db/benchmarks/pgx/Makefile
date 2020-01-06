db: ## Run new container with PostgreSQL.
	docker-compose down && docker-compose up

fill: ## Fill db with random data.
	go run main.go

bench: ## Run benchmark.
	go test -bench=. benchmark_test.go main.go pgx.go sqlx.go gorm.go

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo " 1) run container with 'make db' \n 2) fill db with 'make fill' \n 3) run benchmark with 'make bench'"
.DEFAULT_GOAL := help

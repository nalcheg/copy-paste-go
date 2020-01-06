compile: ## Compile
	CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -a -installsuffix cgo `go-file`

coverage: ## Generate and open in browser tests coverage
	go test -coverprofile=coverage.out -v ./...
	go tool cover -html=coverage.out
	rm coverage.out

govet: ## Run go vet tool
	go vet ./...

golangci-lint: ## Run https://golangci.com/
	docker run -ti --rm -v `pwd`:/goapp golangci/build-runner golangci-lint -v run

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

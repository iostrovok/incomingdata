# Include go binaries into path
export PATH := $(GOPATH)/bin:$(PATH)

mod-action-%:
	@echo "Run go mod ${*}...."
	GO111MODULE=on go mod $*
	@echo "Done go mod  ${*}"

mod: mod-action-verify mod-action-tidy mod-action-vendor mod-action-download mod-action-verify ## Download all dependencies
	@echo "All installed"


test: mod ## Run test
	PRELOAD_BY_VERSION=0 go test -race ./...


tests-cover: ## Run test covering
	PRELOAD_BY_VERSION=1 go test -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	rm coverage.out

clean_cache:
	@go clean -cache
	@go clean -testcache

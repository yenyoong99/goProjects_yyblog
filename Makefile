PKG := "github.com/yenyoong99/goProjects_yyblog"

dep: ## Get the dependencies
	@go mod tidy

run: ## Run Server
	@go run main.go start

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

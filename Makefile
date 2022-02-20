.PHONY:

test:
	go test -coverpkg=./... -coverprofile=coverage.out ./...
	go tool cover -func coverage.out

lint:
	golangci-lint run

lint.fix: ## Lint
	golangci-lint run --fix

generate: ## Generate
	go generate ./...

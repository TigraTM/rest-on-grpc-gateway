.PHONY:

# Unit test
test:
	go test -coverpkg=./... -coverprofile=coverage.out ./...
	go tool cover -func coverage.out

# Lint
lint:
	golangci-lint run

# Lint with flag fix
lint.fix:
	golangci-lint run --fix

# Generate
generate:
	go generate ./...

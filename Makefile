.PHONY:

# Run unit test
test:
	go test -coverpkg=./... -coverprofile=coverage.out ./...
	go tool cover -func coverage.out

# Run go lint
go_lint:
	golangci-lint run

# Run go lint with flag fix
go_lint.fix:
	golangci-lint run --fix

# Run buf lint
buf_lint:
	buf lint

# Generate go and buf
generate:
	go generate ./...
	buf generate

migrations.create:
	migrate create -ext sql -dir modules/$(module)/migrations -format 2006010215 $(issue)

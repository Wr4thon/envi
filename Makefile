validate: lint test

install-tools: install-lint-tool

install-lint-tool:
	go install golang.org/x/lint/golint@latest

lint:
	golint ./...

test:
	CGO_ENABLED=0 go test -mod=mod -short -cover -v -coverprofile=coverage.out -covermode=atomic ./...
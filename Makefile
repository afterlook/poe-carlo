GOLANGCI_LINT_VERSION := v2.7.2
GOLANGCI_LINT = go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

GOTESTSUM_VERSION := v1.13.0
GOTESTSUM = go run gotest.tools/gotestsum@$(GOTESTSUM_VERSION)

.PHONY: vet
fmt:
	@go fmt ./...

.PHONY: vet
vet:
	@go vet ./...

.PHONY: lint
lint: vet fmt
	@$(GOLANGCI_LINT) run

.PHONY: build
build: vet fmt
	@go build -o app main.go

.PHONY: test
test: vet fmt
	@$(GOTESTSUM) --format testname -- -race -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: data
data:
	@./hack/pull-poe-data.sh

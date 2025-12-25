GOLANGCI_LINT_VERSION := v2.7.2
GOTESTSUM_VERSION := v1.13.0

LOCALBIN ?= $(shell pwd)/bin
$(LOCALBIN):
	mkdir -p "$(LOCALBIN)"

GOLANGCI_LINT = $(LOCALBIN)/golangci-lint
GOTESTSUM = $(LOCALBIN)/gotestsum

.PHONY: golangci-lint
golangci-lint: $(GOLANGCI_LINT)
$(GOLANGCI_LINT): $(LOCALBIN)
	$(call go-install-tool,$(GOLANGCI_LINT),github.com/golangci/golangci-lint/v2/cmd/golangci-lint,$(GOLANGCI_LINT_VERSION))

.PHONY: gotestsum
gotestsum: $(GOTESTSUM)
$(GOTESTSUM): $(LOCALBIN)
	$(call go-install-tool,$(GOTESTSUM),gotest.tools/gotestsum,$(GOTESTSUM_VERSION))

.PHONY: tools
tools: golangci-lint gotestsum

.PHONY: vet
fmt:
	@go fmt ./...

.PHONY: vet
vet:
	@go vet ./...

.PHONY: lint
lint: golangci-lint vet fmt
	@"$(GOLANGCI_LINT)" run

.PHONY: build
build: vet fmt
	@go build -o app main.go

.PHONY: test
test: gotestsum vet fmt
	@"$(GOTESTSUM)" --format testname -- -race -coverprofile=coverage.out -covermode=atomic ./...

.PHONY: data
data:
	@./hack/pull-poe-data.sh

# go-install-tool will 'go install' any package with custom target and name of binary, if it doesn't exist
# NOTE: This function was adapted from the Kubebuilder project (https://github.com/kubernetes-sigs/kubebuilder)
# which is licensed under the Apache License 2.0 (https://www.apache.org/licenses/LICENSE-2.0)
# $1 - target path with name of binary
# $2 - package url which can be installed
# $3 - specific version of package
define go-install-tool
@[ -f "$(1)-$(3)" ] && [ "$$(readlink -- "$(1)" 2>/dev/null)" = "$(1)-$(3)" ] || { \
set -e; \
package=$(2)@$(3) ;\
echo "Downloading $${package}" ;\
rm -f "$(1)" ;\
GOBIN="$(LOCALBIN)" go install $${package} ;\
mv "$(LOCALBIN)/$$(basename "$(1)")" "$(1)-$(3)" ;\
} ;\
ln -sf "$$(realpath "$(1)-$(3)")" "$(1)"
endef

.PHONY: vet
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...

.PHONY: lint
lint: vet fmt
	golangci-lint run

.PHONY: build
build: vet fmt
	go build -o app main.go

.PHONY: test
test: vet fmt
	go test ./... --shuffle=on -race -coverprofile=coverage.out -covermode=atomic

.PHONY: data
data:
	./hack/pull-poe-data.sh

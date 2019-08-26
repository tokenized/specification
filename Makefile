# command to build and run on the local OS.
GO_BUILD = go build

# tools
BINARY_CONTRACT_CLI=tokenized

GO_DIST_DIR=dist/golang

all: prepare tools run-generate format test

run-win: prepare-win tools run-generate format-win

protobuf:
	protoc --proto_path=dist/protobuf --go_out=plugins=grpc:dist/golang/actions --js_out=library=actions,binary:dist/typescript/protobuf/actions dist/protobuf/actions.proto
	protoc --proto_path=dist/protobuf --go_out=plugins=grpc:dist/golang/assets --js_out=library=assets,binary:dist/typescript/protobuf/assets dist/protobuf/assets.proto
	protoc --proto_path=dist/protobuf --go_out=plugins=grpc:dist/golang/messages --js_out=library=messages,binary:dist/typescript/protobuf/messages dist/protobuf/messages.proto

generate-code:
	go run cmd/$(BINARY_CONTRACT_CLI)/main.go generate
	goimports -w $(GO_DIST_DIR)

run-generate: generate-code protobuf
	goimports -w $(GO_DIST_DIR)

dist-cli:
	$(GO_DIST) -o dist/$(BINARY_CONTRACT_CLI) cmd/$(BINARY_CONTRACT_CLI)/main.go

format: format-go

lint: lint-go

test: test-go

format-go:
	goimports -w $(GO_DIST_DIR)

lint-go:
	golint $(GO_DIST_DIR)
	go vet ./...

# run the tests with coverage
test-go:
	cd $(GO_DIST_DIR) && go test ./...

tools:
	go get golang.org/x/lint/golint
	go get golang.org/x/tools/cmd/goimports

prepare:
	mkdir -p tmp

format-win:
	goimports -w dist\golang\protocol\

prepare-win:
	mkdir tmp | echo tmp exists

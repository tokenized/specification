FILENAME ?= tmp/protocol-working.csv
FILENAME_ASSET_TYPES ?= tmp/asset_types.json
PACKAGE ?= protocol

# command to build and run on the local OS.
GO_BUILD = go build

ci: clean tools prepare deps test lint

clean:

tools:
	go get github.com/golang/lint/golint
	go get golang.org/x/tools/cmd/goimports

prepare:
	mkdir -p dist

deps:
	go get ./...

utils: tokenized-decode-message tokenized-decode-tx utils-install

utils-install:
	cp dist/tokenized-decode* $(GOPATH)/bin/

generate-go: generate-go-code \
		generate-go-asset-types-code \
		generate-go-forms \
		generate-go-asset-type-forms \
		generate-go-tests lint

generate-go-code:
	go run cmd/tokenized-code-gen/main.go -filename $(FILENAME) -package $(PACKAGE) | goimports > messages.go

generate-go-asset-types-code:
	go run cmd/tokenized-asset-types-gen/main.go -filename $(FILENAME_ASSET_TYPES) | goimports > asset_types.go

generate-go-asset-type-forms:
	go run cmd/tokenized-asset-types-gen/main.go -filename $(FILENAME_ASSET_TYPES) -forms true | goimports > asset_type_forms.go

generate-go-forms:
	go run cmd/tokenized-code-gen/main.go -filename $(FILENAME) -package tokenized -forms true | goimports > protocol_forms.go

generate-go-tests:
	go run cmd/tokenized-code-gen/main.go -tests -filename $(FILENAME) -package $(PACKAGE) | goimports > messages_test.go

generate-python:
	go run cmd/tokenized-code-gen/main.go -python -filename $(FILENAME) -package $(PACKAGE) > messages.py

generate-yaml:
	go run cmd/tokenized-code-gen/main.go -yaml -filename $(FILENAME) -package $(PACKAGE) > messages.yaml

# test:
# 	go test

# run the tests with coverage
test:
	go test -coverprofile=tmp/coverage.out

# run tests with coverage and open html file in the browser
#
# See https://blog.golang.org/cover for more output options
test-coverage: test
	go tool cover -html=tmp/coverage.out


lint: golint go-vet

golint:
	golint

go-vet:
	go vet

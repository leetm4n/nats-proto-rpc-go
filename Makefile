ROOT_DIR = $(shell pwd)
COLOR := "\e[1;36m%s\e[0m\n"
PROTO_ROOT := $(ROOT_DIR)/api/proto
PROTO_EXAMPLE_ROOT := $(ROOT_DIR)/examples/proto
PATH := $(PATH):$(PWD)/bin

##### Build tools #####
.PHONY: install-build-tools
install-build-tools:
	@printf $(COLOR) "Installing build tools..."
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest

##### Proto to go generation #####
.PHONY: codegen-proto
codegen-proto:
	@printf $(COLOR) "Generating proto..."
	@(cd $(PROTO_ROOT) && buf generate)

##### Example proto to go generation #####
.PHONY: codegen-proto-example
codegen-proto-example:
	@printf $(COLOR) "Generating proto..."
	@(cd $(PROTO_EXAMPLE_ROOT) && buf generate)

.PHONY: codegen-proto-example-local
codegen-proto-example-local:
	make build-codegenerator
	@printf $(COLOR) "Generating proto..."
	@(cd $(PROTO_EXAMPLE_ROOT) && buf generate --template buf.gen.local.yaml)

##### Build protoc-gem-natsrpcgo #####
.PHONY: build-codegenerator
build-codegenerator:
	@printf $(COLOR) "Building generator plugin..."
	@go build -o bin/ ./cmd/protoc-gen-natsrpcgo

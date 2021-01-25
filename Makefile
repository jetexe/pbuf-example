PROTOFILES = api/messages/v1/message.proto api/messages/v1/text.proto api/services/v1/api.proto api/state/v1/state.proto
BUF_IMAGE = build/proto.bin

.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help
	@printf "\033[33m%s:\033[0m\n" 'Available commands'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[32m%-11s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.SILENT: $(BUF_IMAGE)
$(BUF_IMAGE): $(PROTOFILES) ## Build models from protofiles.
	echo "> Building with buf"
	mkdir -p "build"
	buf build -o $(BUF_IMAGE)

.SILENT: buf-lint
buf-lint: $(BUF_IMAGE) ## Lint proto files.
	echo "> Linting with buf"
	buf lint $(BUF_IMAGE)

.SILENT: buf-breaking
buf-breaking: $(BUF_IMAGE) ## Check breaking changes.
	echo "> Check breaking changes"
	buf breaking --against '.git#branch=main' $(BUF_IMAGE)

.SILENT: buf-generate
buf-generate: buf-lint buf-breaking ## Build models with protoc.
	echo "> Building with protoc"
	buf generate $(BUF_IMAGE)

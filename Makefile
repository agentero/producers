.PHONY: go-bindings
go-bindings: ## Generates go bindings
	protoc -I ./proto \
  		--go_out ./go --go_opt paths=source_relative \
 		--go-grpc_out ./go --go-grpc_opt paths=source_relative \
  		--grpc-gateway_out ./go --grpc-gateway_opt paths=source_relative --grpc-gateway_opt grpc_api_configuration=./proto/rest-annotations.yaml \
	  	proto/producer.proto

.PHONY: markdown-docs
markdown-docs: ## Generates markdown docs
	protoc -I ./proto \
		--doc_opt=./proto/grpc_docs.template,grpc.md \
		--doc_out='./docs' \
		proto/producer.proto

.PHONY: generate-all
generate-all: go-bindings markdown-docs

.PHONY: help
help: ## Shows help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
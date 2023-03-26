proto: ## Generate protobuf 
	@printf "generate protobuf files \n";
	@for f in grpc/proto/*.proto grpc/proto/*.proto; do \
		protoc --go_out=. $$f; \
		protoc --go-grpc_out=. $$f; \
		echo compiled: $$f; \
	done

	@printf "\n";
	@printf "inject tags \n";

	@for f in grpc/pb/*.go; do \
		protoc-go-inject-tag -input=$$f; \
		echo injected tags: $$f; \
	done
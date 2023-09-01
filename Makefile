generate:
	make generate-segment-api && make generate-user-api

.PHONY: generate-segment-api
generate-segment-api:
	mkdir -p pkg/segment_api
	protoc -I api/segment \
	-I proto --go_out=pkg/segment_api \
	--go_opt=paths=import --go-grpc_out=pkg/segment_api \
	--go-grpc_opt=paths=import --grpc-gateway_out=pkg/segment_api \
	--grpc-gateway_opt=logtostderr=true \
	--grpc-gateway_opt=paths=import api/segment/segment.proto \
	--validate_out lang=go:pkg/segment_api \
	--swagger_out=allow_merge=true,merge_file_name=api:pkg/segment_api 
	mv pkg/segment_api/github.com/nikitads9/segment-service-api/pkg/segment_api/* pkg/segment_api/
	rm -r  ./pkg/segment_api/github.com

.PHONY: generate-user-api
generate-user-api:
	mkdir -p pkg/user_api
	protoc -I api/user \
	-I proto --go_out=pkg/user_api \
	--go_opt=paths=import --go-grpc_out=pkg/user_api \
	--go-grpc_opt=paths=import --grpc-gateway_out=pkg/user_api \
	--grpc-gateway_opt=logtostderr=true \
	--grpc-gateway_opt=paths=import api/user/user.proto \
	--validate_out lang=go:pkg/user_api \
	--swagger_out=allow_merge=true,merge_file_name=api:pkg/user_api 
	mv pkg/user_api/github.com/nikitads9/segment-service-api/pkg/user_api/* pkg/user_api/
	rm -r  ./pkg/user_api/github.com


PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		mkdir -p proto
		cp api/segment/segment.proto proto/
		@if [ ! -d proto/google ]; then \
			git clone https://github.com/googleapis/googleapis proto/googleapis &&\
			mkdir -p  proto/google/ &&\
			mv proto/googleapis/google/api proto/google &&\
			rm -rf proto/googleapis ;\
		fi
		@if [ ! -d proto/github.com/envoyproxy ]; then \
			mkdir -p proto/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate proto/protoc-gen-validate &&\
			mv proto/protoc-gen-validate/validate/*.proto proto/validate &&\
			rm -r -f proto/protoc-gen-validate ;\
		fi
		@if [ ! -d proto/google/protobuf ]; then\
			git clone https://github.com/protocolbuffers/protobuf proto/protobuf &&\
			mkdir -p  proto/google/protobuf &&\
			mv proto/protobuf/src/google/protobuf/*.proto proto/google/protobuf &&\
			rm -rf proto/protobuf ;\
		fi

.PHONY: deps
deps: install-go-deps

.PHONY: install-go-deps
install-go-deps: .install-go-deps

.PHONY: .install-go-deps
.install-go-deps:
		ls go.mod || go mod init
			go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
			go get -u github.com/golang/protobuf/proto
			go get -u github.com/golang/protobuf/protoc-gen-go
			go get -u github.com/envoyproxy/protoc-gen-validate
			go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
			go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
			go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
			go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
			go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
			go install github.com/envoyproxy/protoc-gen-validate
			go install github.com/golang/protobuf/protoc-gen-go
			go install github.com/golang/protobuf/proto
			go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
			go install -v golang.org/x/tools/gopls@latest
			go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
			go mod tidy
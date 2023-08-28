.PHONY: generate
generate:
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
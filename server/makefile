API_PROTO_FILES=$(shell find ../protobuf/api -name *.proto)

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
wire:
	wire ./biz/service

.PHONY: api
api:
	protoc --proto_path=../protobuf/api \
    		--proto_path=../protobuf/third_party \
    		--go_out=paths=source_relative:./gen/proto \
    		--go-http_out=paths=source_relative:./gen/proto \
    		--openapi_out=fq_schema_naming=true,default_response=false:../protobuf/ \
    		 $(API_PROTO_FILES)

.PHONY: dao
dao:
	go run cmd/init_dao/main.go

.PHONY: table
table:
	go run cmd/init_table/*
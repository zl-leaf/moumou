API_PROTO_FILES=$(shell find ../protobuf/api -name *.proto)

wire:
	wire ./biz/service

api:
	protoc --proto_path=../protobuf/api \
    		--proto_path=../protobuf/third_party \
    		--go_out=paths=source_relative:./proto_gen \
    		--go-http_out=paths=source_relative:./proto_gen \
    		 $(API_PROTO_FILES)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: server
server:
	protoc --proto_path=./api \
		--proto_path=./third_party \
		--go_out=paths=source_relative:../server/proto_gen \
		--go-http_out=paths=source_relative:../server/proto_gen \
		 $(API_PROTO_FILES)
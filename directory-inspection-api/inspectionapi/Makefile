ifeq ($(GOPATH),)
	GOPATH := $(shell go env GOPATH)
endif

.PHONY: all
all: generate

.PHONY: dependencies
dependencies:
	GO111MODULE=off go get \
		github.com/gogo/protobuf/gogoproto \
		github.com/gogo/protobuf/protoc-gen-gogofast \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

.PHONY: generate
generate: clean
	@# Generate go, gateway, validators
	protoc -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		*.proto

	@# Generate swagger.json
	protoc -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		inspectionapi.proto
	echo 'package inspectionapi'              >  inspectionapi.swagger.go
	echo 'const ('                            >> inspectionapi.swagger.go
	echo 'SwaggerJSONDirectoryInspection = `' >> inspectionapi.swagger.go
	cat inspectionapi.swagger.json            >> inspectionapi.swagger.go
	echo '`)'                                 >> inspectionapi.swagger.go

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f *.pb.go
	rm -f *.pb.gw.go
	rm -f *.swagger.json
	rm -f *.swagger.go

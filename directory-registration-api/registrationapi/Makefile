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

protos = $(wildcard *.proto)
results = $(protos:.proto=.pb.go) $(protos:.proto=.pb.gw.go) $(protos:.proto=.swagger.json) $(protos:.proto=.swagger.go)

%.pb.go %.pb.gw.go: %.proto
	protoc -I. \
		-I$(shell go env GOPATH)/src \
		-I$(shell go env GOPATH)/src/github.com/gogo/protobuf/protobuf \
		-I$(shell go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogofast_out=\
Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,\
plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		$<

%.swagger.json: %.proto
	@# Generate swagger.json
	protoc -I. \
		-I$(shell go env GOPATH)/src \
		-I$(shell go env GOPATH)/src/github.com/gogo/protobuf/protobuf \
		-I$(shell go env GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:. \
		$<

%.swagger.go: %.swagger.json
	echo 'package $*'               >  $@
	echo 'const ('                  >> $@
	echo 'SwaggerJSONDirectory = `' >> $@
	cat $<                          >> $@
	echo '`)'                       >> $@

.PHONY: generate
generate: dependencies $(results)

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f *.pb.go
	rm -f *.pb.gw.go
	rm -f *.swagger.json
	rm -f *.swagger.go

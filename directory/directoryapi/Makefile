.PHONY: all
all: generate

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
		directory.proto
	echo 'package directoryapi'      > directory.swagger.go
	echo 'const ('                  >> directory.swagger.go
	echo 'SwaggerJSONDirectory = `' >> directory.swagger.go
	cat directory.swagger.json      >> directory.swagger.go
	echo '`)'                       >> directory.swagger.go

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f *.pb.go
	rm -f *.pb.gw.go
	rm -f *.swagger.json
	rm -f *.swagger.go

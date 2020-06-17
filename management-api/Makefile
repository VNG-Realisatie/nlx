ifeq ($(GOPATH),)
	GOPATH := $(shell go env GOPATH)
endif

API_DIR := $(PWD)/api
PKG_DIR := $(PWD)/pkg

PROTO_NAMES:=configapi directory

PROTO_MOCK_FILES := $(foreach proto,$(PROTO_NAMES),$(PKG_DIR)/$(proto)/mock/mock_$(proto).go)
PROTO_GO_FILES := $(foreach proto,$(PROTO_NAMES),$(PKG_DIR)/$(proto)/$(proto).pb.go)
PROTO_GATEWAY_FILES := $(foreach proto,$(PROTO_NAMES),$(PKG_DIR)/$(proto)/$(proto).pb.gw.go)
PROTO_SWAGGER_FILES := $(PROTO_NAMES:%=$(API_DIR)/%.swagger.json)

EXTRA_MOCK_FILES := $(PKG_DIR)/configapi/mock/mock_database.go

GENERATE_FILES := $(PROTO_GO_FILES) $(PROTO_SWAGGER_FILES) $(PROTO_MOCK_FILES) $(EXTRA_MOCK_FILES)

.PHONY: all
all: clean generate

.PHONY: generate
generate: $(GENERATE_FILES)

.PHONY: dependencies
dependencies:
	GO111MODULE=off go get \
		github.com/gogo/protobuf/gogoproto \
		github.com/gogo/protobuf/protoc-gen-gogofast \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
		github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen

# A template is used instead of a pattern rule because the name is used in the directory and filename.
define PROTOBUFFER_RULES

$(PKG_DIR)/$(PROTO)/$(PROTO).pb.go $(PKG_DIR)/$(PROTO)/$(PROTO).pb.gw.go: $(API_DIR)/$(PROTO).proto
	@# Generate go, gateway
	cd $$(<D) && protoc -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--gogofast_out=$$\
			Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,$$\
			Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,$$\
			Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,$$\
			Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,$$\
			Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,$$\
			plugins=grpc:$$(@D) \
		--grpc-gateway_out=logtostderr=true:$$(@D) \
		$$(<F)

$(PKG_DIR)/$(PROTO)/mock/mock_$(PROTO).go: $(PKG_DIR)/$(PROTO)/$(PROTO).pb.go
	@# Generate mock file
	cd $$(@D) && mockgen -source $$(shell realpath --relative-to=$$(@D) $$<) -destination $$(@F)

endef # PROTOBUFFER_RULES

$(foreach PROTO,$(PROTO_NAMES),$(eval $(PROTOBUFFER_RULES)))

$(PKG_DIR)/configapi/mock/mock_database.go: $(PKG_DIR)/configapi/database.go
	@# Generate mock file
	cd $(@D) && mockgen -source $(shell realpath --relative-to=$(@D) $<) -destination $(@F)

$(API_DIR)/%.swagger.json: $(API_DIR)/%.proto
	@# Generate swagger.json
	cd $(<D) && protoc -I. \
		-I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/gogo/protobuf/protobuf \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:$(@D) \
		$(<F)

.PHONY: clean
clean:
	@# Remove old generated files
	rm -f $(GENERATE_FILES)

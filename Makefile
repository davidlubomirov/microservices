GO_PROTO_DIR=${PWD}

build-go-proto-auth:
	protoc proto/auth.proto -I. --go_out=${GO_PROTO_DIR}


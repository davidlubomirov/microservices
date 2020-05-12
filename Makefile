GO_PROTO_DIR=${PWD}

build-auth-go:
	protoc -I. --go_out=plugins=grpc:proxy proto/auth.proto

build-python-proto-auth:
	#protoc proto/auth.proto -I. --python_out=${GO_PROTO_DIR}/auth
	python -m grpc_tools.protoc -I. --python_out=${GO_PROTO_DIR}/auth --grpc_python_out=${GO_PROTO_DIR}/auth proto/auth.proto

SHELL := /usr/bin/bash
.SILENT: run

GO_OUT_PATH = internal/adapters/framework/left/grpc
PROTO_PATH = ${GO_OUT_PATH}/proto/

protogen:
	@protoc --go_out=${GO_OUT_PATH} --proto_path=${PROTO_PATH} ${PROTO_PATH}number_msg.proto
	@protoc --go_grpc_out=${GO_OUT_PATH} --proto_path=${PROTO_PATH} ${PROTO_PATH}arithmetic_svc.proto
	@echo "Generated at ${PROTO_PATH}"

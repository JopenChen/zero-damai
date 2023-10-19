
.PHONY: gen-all
gen-all:
	make gen-api gen-rpc

.PHONY: gen-api
gen-api:
	goctl api go --api .\gateway\api\gateway.api --dir .\gateway\ --style=go_zero
	@echo "Generate gateway.api file successfully"

.PHONY: gen-rpc
gen-rpc:
	goctl rpc protoc service\proto\service.proto --go_out=./service --go-grpc_out=./service --zrpc_out=./service --style=go_zero -m -c=false
	@echo "Generate service.proto file successfully"
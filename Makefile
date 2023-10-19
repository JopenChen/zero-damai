
.PHONY: gen-api
gen-api:
	goctl api go --api .\gateway\api\gateway.api --dir .\gateway\ --style=go_zero
	@echo "Generate gateway.api file successfully"

.PHONY: gen-rpc
gen-rpc:
	goctl api go --api .\gateway\api\gateway.api --dir .\gateway\ --style=go_zero
	@echo "Generate gateway.api file successfully"

.PHONY: gen-all
gen-all:
	make gen-api gen-rpc gen-model

.PHONY: gen-api
gen-api:
	goctl api go --api=.\gateway\api\gateway.api --dir=.\gateway --style=go_zero --home=.\template\1.5.6
	@echo "Generate gateway.api file successfully"

.PHONY: gen-rpc
gen-rpc:
	goctl rpc protoc service\proto\service.proto --go_out=./service --go-grpc_out=./service --zrpc_out=./service --style=go_zero -m --home .\template\1.5.6
	@echo "Generate service.proto file successfully"

.PHONY: gen-model
gen-model:
	goctl model mysql ddl --src .\database\mysql\user.sql --dir .\model\mysql\user --style go_zero --home .\template\1.5.6
	@echo "Generate user.sql model file successfully"

	goctl model mysql ddl --src .\database\mysql\performance.sql --dir .\model\mysql\performance --style go_zero --home .\template\1.5.6
	@echo "Generate performance.sql model file successfully"

	goctl model mysql ddl --src .\database\mysql\performance_session.sql --dir .\model\mysql\performance_session --style go_zero --home .\template\1.5.6
	@echo "Generate performance_session.sql model file successfully"

	goctl model mysql ddl --src .\database\mysql\performance_seat.sql --dir .\model\mysql\performance_seat --style go_zero --home .\template\1.5.6
	@echo "Generate performance_seat.sql model file successfully"

	#goctl model mysql ddl --src .\database\mysql\order.sql --dir .\model\mysql\order --style go_zero --home .\template\1.5.6\
#    @echo "Generate order.sql model file successfully"
	#goctl model mysql ddl --src .\database\mysql\log.sql --dir .\model\mysql\log --style go_zero --home .\template\1.5.6\
#    @echo "Generate log.sql model file successfully"






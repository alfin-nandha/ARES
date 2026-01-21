PROTO_PATH=proto/*.proto
APP_NAME=ARES

gen-grpc:
	@echo "generating go file......"
	@protoc --go_out=. --go-grpc_out=. ${PROTO_PATH}

run:
	@echo "running ${APP_NAME}......"
	@go run main.go 

test:
	@echo "running all test file......"
	@go test -v ./test/..
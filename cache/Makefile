build:
	go build -o ./build/cache_manager ./cmd/cache_manager

test:
	go test ./...

proto-generate:
	protoc --go_out=./ --go-grpc_out=./ proto/cache_manager.proto
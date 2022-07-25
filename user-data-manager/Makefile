build:
	go build -o ./build/user-data-manager ./cmd/user-data-manager

test:
	go test ./...

proto-generate:
	protoc --go_out=./ --go-grpc_out=./ proto/user-data-manager.proto
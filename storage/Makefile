build:
	go build -o ./build/storage ./cmd/storage

test:
	go test ./...

proto-generate:
	protoc --go_out=./ --go-grpc_out=./ proto/storage.proto
build:
	go build -o ./build/user_authorization ./cmd/user_authorization

test:
	go test ./...

proto-generate:
	protoc --go_out=./ --go-grpc_out=./ proto/user_authorization.proto
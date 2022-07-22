up:
	docker-compose --file docker-compose.yml up --build

down:
	docker-compose --file docker-compose.yml down

restart:
	docker-compose --file docker-compose.yml restart

recreate:
	docker-compose --file docker-compose.yml down
	docker-compose --file docker-compose.yml up -d --build --force-recreate
	docker-compose --file docker-compose.yml up -d

logs:
	docker container logs user_data_manager_go

build:
	go build -o ./build/user-data-manager ./cmd/user-data-manager

exec:
	docker exec -it url_carver_go sh

test:
	go test ./...

proto-generate:
	protoc --go_out=./ --go-grpc_out=./ proto/user-data-manager.proto
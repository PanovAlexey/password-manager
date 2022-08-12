up:
	sh -c ".build/run.sh"
down:
	sh -c ".build/down.sh"
restart:
	sh -c ".build/restart.sh"
recreate:
	sh -c ".build/recreate.sh"
logs:
	docker container logs password-manager-api-gw \
		&& docker container logs password-manager-user-data \
		&& docker container logs password-manager-user-auth \
		&& docker container logs password-manager-storage \
		&& docker container logs password-manager-storage-master-postgres
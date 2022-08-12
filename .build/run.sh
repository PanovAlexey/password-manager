cp -f database/.env.example database/.env
sh -c "cd ./database && docker-compose up -d"

cp -f storage/.env.example storage/.env
sh -c "cd ./storage && docker-compose up -d"

cp -f user/.env.example user/.env
sh -c "cd ./user && docker-compose up -d"

cp -f user-data-manager/.env.example user-data-manager/.env
sh -c "cd ./user-data-manager && docker-compose up -d"

cp -f api-gw/.env.example api-gw/.env
sh -c "cd ./api-gw && docker-compose up -d"
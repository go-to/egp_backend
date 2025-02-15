# debug time value
time=

# env
DOCKER_CONTAINER_API=api
DOCKER_CONTAINER_DB=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=egp
DB_USER=egp_user
DB_PASS=password

# cmd
up:
	docker compose up -d
down:
	docker compose down
logs-api:
	docker logs -f ${DOCKER_CONTAINER_API}
logs-db:
	docker logs -f ${DOCKER_CONTAINER_DB}
db-migrate-up:
	migrate -path migrations -database 'postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=egp&sslmode=disable' -verbose up
db-migrate-down:
	migrate -path migrations -database 'postgresql://${DB_USER}:${DB_PASS}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=egp&sslmode=disable' -verbose down
run:
	go run main.go
grpcurl-shops:
	grpcurl -plaintext localhost:8080 egp.EgpService.GetShops
clean-branch:
	git switch main
	git branch | grep -v "main" | xargs git branch -d
	git pull
update-gomod:
	go get -u
	go mod tidy
# e.x.) make debug-time time='2025-03-11 19:00:00'
debug-time:
	mysql -u${DB_USER} -p${DB_PASS} -P${DB_PORT} -h${DB_HOST} ${DB_NAME} -e "UPDATE config SET conf_value = '${time}' WHERE conf_name = 'debug_time';"

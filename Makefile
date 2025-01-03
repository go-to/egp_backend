DOCKER_CONTAINER_API=api
DOCKER_CONTAINER_DB=db
DB_HOST=localhost
DB_PORT=3306
DB_NAME=egp
DB_USER=root
DB_PASS=root

up:
	docker compose up -d
down:
	docker compose down
logs-api:
	docker logs -f ${DOCKER_CONTAINER_API}
logs-db:
	docker logs -f ${DOCKER_CONTAINER_DB}
db-migrate-up:
	migrate -path db/migrations -database 'mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}' -verbose up
db-migrate-down:
	migrate --path db/migrations --database 'mysql://${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}' -verbose down
run:
	go run main.go
grpcurl-test:
	grpcurl -plaintext localhost:8080 egp.EgpService.GetShops
clean-branch:
	git switch main && git branch | xargs git branch -d

DOCKER_CONTAINER_API="api"
DOCKER_CONTAINER_DB="db"

up:
	docker compose up -d
down:
	docker compose down
logs-api:
	docker logs -f ${DOCKER_CONTAINER_API}
logs-db:
	docker logs -f ${DOCKER_CONTAINER_DB}
run:
	go run main.go
grpcurl-test:
	grpcurl -plaintext localhost:8080 egp.EgpService.GetShops
clean-branch:
	git switch main && git branch | xargs git branch -d

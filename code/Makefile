dockerup:
	cd .. ; cd docker ; docker compose up -d

dockerdown:
	cd .. ; cd docker ; docker compose down

createdb:
	docker exec -it postgres-db createdb --username=postgres --owner=postgres go_finance                                                                   ─╯

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_finance?sslmode=disable" -verbose down  

sqlcgenerate:
	docker run --rm -v $$(pwd):/src -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...	

swaggercreate:
	swag init --parseDependency --parseInternal --parseDepth 1

server:
	go run main.go	

.PHONE:  createdb  migrateup migratedown dockerup sqlcgenerate server

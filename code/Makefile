dockerup:
	cd .. ; cd docker docker compose up 

dockerdown:
	cd .. ; cd docker docker compose down

createdb:
	docker exec -it postgres-db createdb --username=postgres --owner=postgres gofinance                                                                   ─╯

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/go_finance?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/go_finance?sslmode=disable" -verbose down  

sqlcgenerate:
	docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate

.PHONE:  createdb  migrateup migratedown dockerup sqlcgenerate

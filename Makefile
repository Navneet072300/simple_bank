postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=mypass -d postgres:17-alpine  

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres17 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:mypass@localhost:5432/simple_bank?sslmode=disable" -verbose up


migratedown:
	migrate -path db/migration -database "postgresql://root:mypass@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -source file://db/migration -database "postgresql://root:mypass@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -source file://db/migration -database "postgresql://root:mypass@localhost:5432/simple_bank?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test  -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination mock/store.go github.com/techschool/simplebank/sqlc Store

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis

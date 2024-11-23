include .envrc
.PHONY: go/bash
go/bash:
	@sudo docker run -it --rm -v $(pwd):/go/src/app -p 8080:8080 go-dev-env /bin/bash



.PHONY: run
run:
	@go run ./cmd/web -port=4000 -env=production -db-dsn=${PIN_DB_DSN}


.PHONY: db/psql
db/psql:
	psql ${PIN_DB_DSN}

.PHONY: gobash
gobash:
	@sudo docker run -it --rm -v $(pwd):/go/src/app -w /go/src/app --name go-dev-env -p 8080:8080 golang:1.23 /bin/bash

.PHONY: dockerbuild
dockerbuild:
	@sudo docker build -t go-dev-env .

.PHONY: dockercreateDatabase
dockercreateDatabase:
	@sudo docker run --name postgres-dbthree -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=mydb -p 2222:2222 -d postgres

.PHONY: dockerenterDatabase
dockerenterDatabase:
	@sudo docker exec -it postgres-dbthree psql -U user -d mydb

.PHONY: dockerstartgolang
dockerstartgolang:
	@sudo docker run -it --rm -v $(pwd):/go/src/app -w /go/src/app --network my-network -p 4000:4000 -e PIN_DB_DSN="postgres://user:password@database:5432/mydb?sslmode=disable" go-dev-env go run ./cmd/web -port=4000 -env=production -db-dsn=postgres://user:password@database:5432/mydb?sslmode=disable

.PHONY: startdbs
startdbs:
	@sudo docker start postgres-dbthree

.PHONY: includepostgresinnetwork
includepostgresinnetwork:
	@sudo docker network connect my-network postgres-dbthree

.PHONY: restartdb
restartdb:
	@sudo docker restart postgres-dbthree
.PHONY: rnew
rnew:
	@sudo docker run -it --rm -v $(pwd):/go/src/app -w /go/src/app --network my-network -p 4000:4000 -e PIN_DB_DSN="postgres://user:password@postgres-dbthree:5432/mydb?sslmode=disable" go-dev-env go run ./cmd/web -port=4000 -env=production -db-dsn=postgres://user:password@postgres-dbthree:5432/mydb?sslmode=disable

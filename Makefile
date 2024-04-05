## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run/api: run the cmd/api application
.PHONY: run/api 
run/api:
	air

## db: connect to db using psql
.PHONY: db 
db:
	psql 'postgres://greenlight:pass@127.0.0.1/greenlight?sslmode=disable'

## db/migrations/new name =$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

## db/migrations/up: apply all up migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	migrate -path ./migrations/ -database 'postgres://greenlight:pass@127.0.0.1/greenlight?sslmode=disable' up

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

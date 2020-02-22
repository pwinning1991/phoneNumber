.PHONY: psql-up psql-down psql-connect

psql-up:
	@docker run --rm -d --name psql1 --env POSTGRES_USER \
		--env POSTGRES_PASSWORD \
		--env POSTGRES_DB \
		-p 5432:5432 \
		postgres:9.5

psql-down:
	@docker stop psql1

psql-connect:
	@docker exec -it psql1 psql --username=$(POSTGRES_USER) --dbname=$(POSTGRES_DB) 

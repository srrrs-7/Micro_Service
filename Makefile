init:
	cp compose.yml.sample conpose.override.yml

up:
	docker compose up

down:
	docker compose down

migrate_file:
	migrate create -ext sql -dir migration migrate.sql

migrate_up:
	migrate -path migration -database "mysql://root:secret@tcp(localhost:3306)/aic" up

migrate_down:
	migrate -path migration -database "mysql://root:secret@tcp(localhost:3306)/aic" down


run:
	go run ./src/cmd

time:
	time go run ./src/cmd

up:
	docker compose up -d --build

down:
	docker compose down

build:
	go build ./src/cmd/main.go
	./src/cmd//main

clear_cache:
	go clean -modcache

endpoint:
	grep 'router.\S' src/util/http/serve.go
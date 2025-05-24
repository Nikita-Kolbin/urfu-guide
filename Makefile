swagger:
	swag init -g ./cmd/main.go

build: swagger
	go build -o build.exe ./cmd/main.go

run: swagger
	go run ./cmd/main.go

compose_up:
	docker compose --env-file ./.env up -d

compose_down:
	docker compose down

compose_drop:
	docker compose down -v

compose_rebuild:
	docker compose down -v
	docker compose --env-file ./.env up -d --build --force-recreate


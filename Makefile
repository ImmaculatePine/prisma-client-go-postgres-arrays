up:
	docker compose up -d

down:
	docker compose down

migrate:
	go run github.com/steebchen/prisma-client-go migrate dev --schema ./prisma/schema.prisma

run: migrate
	go run main.go

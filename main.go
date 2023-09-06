package main

import (
	"context"
	"log"

	"github.com/ImmaculatePine/prisma-client-go-postgres-arrays/prisma/db"
)

func main() {
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatal(err)
		}
	}()

	ctx := context.Background()

	// Works
	createUser(ctx, client)

	// Works
	findUsers(ctx, client)

	// Fails with
	// could not send raw query:
	// 		json data result unmarshal:
	// 			json:
	// 				cannot unmarshal object into Go struct field RawUserModel.items of type raw.String
	findUsersRaw(ctx, client)
}

func createUser(ctx context.Context, client *db.PrismaClient) {
	createdUser, err := client.User.CreateOne(
		db.User.Name.Set("Alice"),
		db.User.Items.Set([]string{"a", "b", "c"}),
	).Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created user", createdUser)
}

func findUsers(ctx context.Context, client *db.PrismaClient) {
	users, err := client.User.FindMany().Exec(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Found users", users)
}

func findUsersRaw(ctx context.Context, client *db.PrismaClient) {
	var res []db.RawUserModel

	if err := client.Prisma.QueryRaw(`
		SELECT * FROM "users"
	`).Exec(ctx, &res); err != nil {
		log.Fatal(err)
	}

	log.Println("Found raw users", res)
}

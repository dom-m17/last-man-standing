package main

import (
	"context"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	db "github.com/dom-m17/lms/backend/db/sqlc"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := db.New(conn)

	// create a user
	insertedUser, err := queries.CreateUser(ctx, db.CreateUserParams{
		Username:       "dom_m17",
		HashedPassword: "password",
		FirstName:      "Dominic",
		LastName:       "Maynard",
		Email:          "dom@email.com",
		PhoneNumber:    pgtype.Text{String: "0123", Valid: true},
		FavouriteTeam:  pgtype.Int8{Int64: 1, Valid: true},
	})
	if err != nil {
		return err
	}
	log.Println(insertedUser)

	// get the author we just inserted
	fetchedUser, err := queries.GetUser(ctx, insertedUser.ID)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedUser, fetchedUser))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pg_database "github.com/nt2311-vn/fullstack_go_vue/internal/database/postgres/compile"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot loading environment variable into app", err)
	}

	connStr := os.Getenv("PG_CONNSTR")

	if connStr == "" {
		log.Fatal("pleasae recheck database connection string")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("cannot connect to database service %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("cannot ping the database service %v", err)
	}

	store := pg_database.New(db)
	ctx := context.Background()

	_, err = store.CreateUsers(ctx, pg_database.CreateUsersParams{
		UserName:     "testuser",
		PassWordHash: "hash",
		Name:         "test",
	})
	if err != nil {
		log.Fatalln("error creating user: ", err)
	}

	eid, err := store.CreateExercise(ctx, "Exercise1")
	if err != nil {
		log.Fatalln("error creating exercise ", err)
	}

	set, err := store.CreateSet(ctx, pg_database.CreateSetParams{
		ExerciseID: eid,
		Weight:     100,
	})
	if err != nil {
		log.Fatalln("error updating exercise: ", err)
	}

	set, err = store.UpdateSet(ctx, pg_database.UpdateSetParams{
		ExerciseID: eid,
		SetID:      set.SetID,
		Weight:     200,
	})
	if err != nil {
		log.Fatalln("error updating set: ", err)
	}

	log.Println("Done!")

	u, err := store.ListUsers(ctx)
	if err != nil {
		log.Fatalln("error listing users: ", err)
	}

	for _, user := range u {
		fmt.Printf("Name: %s, ID: %d\n", user.Name, user.UserID)
	}
}

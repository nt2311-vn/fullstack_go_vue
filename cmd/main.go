package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"

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

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("cannot ping the database service %v", err)
	}

	if err = runSchema(db); err != nil {
		log.Fatalln("cannot run executing schema file", err)
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

func runSchema(db *sql.DB) error {
	schemaFile := filepath.Join("internal", "postgres", "schema", "schema.sql")
	schema, err := os.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("cannot read schema file %v", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("cannot execute sql schema file: %v", err)
	}

	return nil
}

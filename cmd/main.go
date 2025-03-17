package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	pg_database "github.com/nt2311-vn/fullstack_go_vue/internal/database/postgres/compile"
	"github.com/nt2311-vn/fullstack_go_vue/internal/logger"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("cannot loading environment variable into app", err)
	}

	connStr := os.Getenv("PG_CONNSTR")

	if connStr == "" {
		log.Fatal("pleasae recheck database connection string")
	}

	l := flag.Bool("local", false, "true - send to stdout, false - send to logging server")
	flag.Parse()
	logger.SetLoggingOutput(*l)

	logger.Logger.Debugf("Application logging to stdout =%v", *l)
	logger.Logger.Info("Starting the application...")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Logger.Errorf("cannot connect to database service: %s", err.Error())
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		logger.Logger.Errorf("cannot ping the database service: %s", err.Error())
	}

	logger.Logger.Info("Database connected")

	store := pg_database.New(db)
	ctx := context.Background()

	_, err = store.CreateUsers(ctx, pg_database.CreateUsersParams{
		UserName:     "testuser2",
		PassWordHash: "hash",
		Name:         "test2",
	})
	if err != nil {
		logger.Logger.Fatalf("error creating user: %v", err)
	}

	logger.Logger.Info("Success - user created!")

	eid, err := store.CreateExercise(ctx, "Exercise1")
	if err != nil {
		logger.Logger.Errorf("error creating exercise: %v", err)
	}

	logger.Logger.Info("Success - exercise created!")

	set, err := store.CreateSet(ctx, pg_database.CreateSetParams{
		ExerciseID: eid,
		Weight:     100,
	})
	if err != nil {
		logger.Logger.Errorf("error creating set: %v", err)
	}

	logger.Logger.Info("Success - sets created!")

	set, err = store.UpdateSet(ctx, pg_database.UpdateSetParams{
		ExerciseID: eid,
		SetID:      set.SetID,
		Weight:     200,
	})
	if err != nil {
		logger.Logger.Errorf("error updating set: %v ", err)
	}

	logger.Logger.Info("Success - sets updated!")
	logger.Logger.Info("Application complete!")

	defer time.Sleep(1 * time.Second)
}

func runSchema(db *sql.DB) error {
	schemaFile := filepath.Join("internal", "database", "postgres", "schema", "schema.sql")
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

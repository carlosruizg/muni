package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/carlosruizg/muni"
	"github.com/carlosruizg/muni/ent/migrate"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// Create ent.Client and run the schema migration.
	dbClient, err := muni.NewDB()
	defer func() {
		err = errors.Join(err, dbClient.CloseDB())
		fmt.Printf("ERROR: %v", err)
	}()

	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := dbClient.EntClient.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "test_data" {
		createTestData(dbClient)
	}
	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(muni.NewSchema(dbClient.EntClient))
	http.Handle("/",
		playground.Handler("Expert", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}

func createTestData(dbClient *muni.DB) {
	// dbClient.EntClient.Qualification.Create().set
}

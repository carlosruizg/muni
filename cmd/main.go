package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

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

	// client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := dbClient.EntClient.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
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

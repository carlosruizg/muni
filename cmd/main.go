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
	ctx := context.Background()
	if err := dbClient.EntClient.Schema.Create(
		ctx,
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	if len(os.Args) > 1 && os.Args[1] == "test_data" {
		createTestData(ctx, dbClient)
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

func createTestData(ctx context.Context, dbClient *muni.DB) {
	experts, _ := dbClient.EntClient.Expert.Query().All(ctx)
	fmt.Printf("%+v \n", experts)
	quali, _ := dbClient.EntClient.Qualification.Query().All(ctx)
	fmt.Printf("%+v \n", quali)
	// quali, + := dbClient.EntClient.
	// dbClient.EntClient.Qualification.Create().SetValue(schema.Coder).SaveX(ctx)
	// dbClient.EntClient.Qualification.Create().SetValue(schema.Medical).SaveX(ctx)
}

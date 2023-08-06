package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/carlosruizg/muni/ent"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const defaultPort = "8080"

func main() {
	client := Open("postgresql://carlosrg:@127.0.0.1/muni")
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// ctx := context.Background()
}

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func CreateExpert(ctx context.Context, client *ent.Client) (*ent.Expert, error) {
	u, err := client.Expert.
		Create().
		SetName("Carlos").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

// func main() {
// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = defaultPort
// 	}

// 	repo := NewTaskRepo()
// 	srv := handler.NewDefaultServer(
// 		graph.NewExecutableSchema(
// 			graph.Config{
// 				Resolvers: &graph.Resolver{
// 					LabellingTask: repo.Tasks,
// 				},
// 			}))

// 	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
// 	http.Handle("/query", srv)

// 	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))
// }

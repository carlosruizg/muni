package muni

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/carlosruizg/muni/ent"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type DB struct {
	EntClient *ent.Client
}

func NewDB() (*DB, error) {
	dbPassword := os.Getenv("MUNI_DATABASE_PASSWORD")
	dbUrl := fmt.Sprintf("postgresql://muni_app:%s@127.0.0.1/muni", dbPassword)
	client := open(dbUrl)
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return &DB{
		EntClient: client,
	}, nil
}

func (db *DB) CloseDB() error {
	return db.EntClient.Close()
}

// Open new connection
func open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

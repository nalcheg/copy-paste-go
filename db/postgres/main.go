package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/tern/migrate"
	"github.com/jmoiron/sqlx"
)

const dbDSN = "user=postgres password=postgres host=127.0.0.1 port=55432 dbname=test sslmode=disable"

func main() {
	db, err := databaseConnect(dbDSN)
	if err != nil {
		panic(err)
	}

	if err := databaseMigrate(db); err != nil {
		panic(err)
	}

	db.Close()

	dbx, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := dbx.Close(); err != nil {
			log.Print(err)
		}
	}()

	s := &ExampleStructWithDb{db: dbx}

	if err := s.ExecuteAll(); err != nil {
		panic(err)
	}
}

func databaseMigrate(db *pgxpool.Pool) error {
	c, err := db.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer c.Release()

	migrator, err := migrate.NewMigrator(context.Background(), c.Conn(), "schema_version")
	if err != nil {
		return err
	}

	if err := migrator.LoadMigrations("migrations"); err != nil {
		return err
	}

	if err := migrator.Migrate(context.Background()); err != nil {
		return err
	}

	return nil
}

func databaseConnect(dbDSN string) (*pgxpool.Pool, error) {
	c, err := pgxpool.ParseConfig(dbDSN)
	if err != nil {
		return nil, err
	}

	c.MaxConns = 5

	//db, err := pgxpool.Connect(context.Background(), dbDSN)
	db, err := pgxpool.ConnectConfig(context.Background(), c)
	if err != nil {
		return nil, err
	}

	return db, nil
}

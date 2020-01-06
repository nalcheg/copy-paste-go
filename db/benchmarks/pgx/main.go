package main

import (
	"context"
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jackc/tern/migrate"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

var dbDSN = "user=postgres password=postgres host=127.0.0.1 port=55432 dbname=test sslmode=disable"

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
			log.Fatal(err)
		}
	}()

	if err := fillDb(dbx); err != nil {
		log.Fatal(err)
	}
}

func fillDb(dbx *sqlx.DB) error {
	gofakeit.Seed(0)
	for i := 0; i < 10; i++ {
		if _, err := dbx.Exec(`INSERT INTO users (name) VALUES ($1)`, gofakeit.Username()); err != nil {
			return err
		}
	}

	for i := 0; i < 100; i++ {
		if _, err := dbx.Exec(
			`INSERT INTO tickets (creator_id, subject, created_at) VALUES ($1, $2, $3)`,
			gofakeit.Number(1, 10),
			gofakeit.HackerPhrase(),
			time.Now(),
		); err != nil {
			return err
		}
	}

	for i := 1; i <= 100; i++ {
		for j := 0; j < 10; j++ {
			if _, err := dbx.Exec(
				`INSERT INTO comments (ticket_id, user_id, comment, created_at) VALUES ($1, $2, $3, $4)`,
				i,
				gofakeit.Number(1, 10),
				gofakeit.HackerPhrase(),
				time.Now(),
			); err != nil {
				return err
			}
		}
	}

	return nil
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

func connectPgx() (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), dbDSN)
}

func connectSqlxWithPgx() (*sqlx.DB, error) {
	return sqlx.Open("pgx", dbDSN)
}

func connectGorm() (db *gorm.DB, err error) {
	return gorm.Open("postgres", dbDSN)
}

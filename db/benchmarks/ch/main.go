package main

import (
	"log"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/mailru/dbr"
	_ "github.com/mailru/go-clickhouse"
)

func main() {
	c, err := dbr.Open("clickhouse", "http://127.0.0.1:8123/?debug=0", nil)
	if err != nil {
		log.Fatal(err)
	}

	if err := migr(c); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}

	if err := fill(c); err != nil {
		log.Fatal(err)
	}
}

func migr(c *dbr.Connection) error {
	driver, err := clickhouse.WithInstance(c.DB, &clickhouse.Config{
		MultiStatementEnabled: true,
	})

	m, err := migrate.NewWithDatabaseInstance(
		//"file://db/benchmarks/ch/migrations",
		"file://migrations",
		"clickhouse", driver)
	if err != nil {
		return err
	}

	return m.Up()
}

func fill(c *dbr.Connection) error {
	gofakeit.Seed(0)
	sess := c.NewSession(nil)
	query := `
		INSERT INTO default.events (
			Id, EventDate, Data0, Data1, Data2, Data3, Data4, Data5, Data6, Data7, Data8, Data9, Data10, Data11, Data12, Data13, Data14, Data15, Data16, Data17, Data18, Data19, Data20, Data21, Data22, Data23, Data24, Data25, Data26, Data27, Data28, Data29
		) VALUES (
			?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?
		)
	`
	for i := 0; i < 1000; i++ {
		var date string
		now := time.Now()
		if i < 200 {
			date = now.Format("2006-01-02")
		} else if i >= 200 && i < 400 {
			date = now.AddDate(0, 0, -1).Format("2006-01-02")
		} else if i >= 400 && i < 600 {
			date = now.AddDate(0, 0, -2).Format("2006-01-02")
		} else if i >= 600 && i < 800 {
			date = now.AddDate(0, 0, -3).Format("2006-01-02")
		} else if i >= 800 && i < 1000 {
			date = now.AddDate(0, 0, -4).Format("2006-01-02")
		}

		if _, err := sess.Exec(
			query,
			i,
			date,
			gofakeit.BeerName(),
			gofakeit.BeerAlcohol(),
			gofakeit.BeerBlg(),
			gofakeit.BeerMalt(),
			gofakeit.BeerStyle(),
			gofakeit.BeerIbu(),
			gofakeit.BeerYeast(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
			gofakeit.BeerName(),
		); err != nil {
			return err
		}
	}

	return nil
}

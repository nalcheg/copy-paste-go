package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type ExampleStructWithDb struct {
	db *sqlx.DB
}

func (e *ExampleStructWithDb) ExecuteAll() error {
	if err := e.Insert(); err != nil {
		return err
	}

	if err := e.SelectViaQueryx(); err != nil {
		return err
	}

	if err := e.SelectViaScan(); err != nil {
		return err
	}

	if err := e.SelectViaSelect(); err != nil {
		return err
	}

	if _, err := e.db.Exec(`TRUNCATE TABLE t1`); err != nil {
		return err
	}

	return nil
}

func (e *ExampleStructWithDb) Insert() error {
	if _, err := e.db.Exec(
		`INSERT INTO t1 (placebo) VALUES ($1), ($2), ($3), ($4), ($5), ($6)`,
		"placebo",
		"placeTwo",
		"placeThree",
		"placeFour",
		"placeFive",
		"placeSix",
	); err != nil {
		return err
	}

	return nil
}

type Row struct {
	ID      uint64 `db:"id"`
	Placebo string `db:"placebo"`
}

func (e *ExampleStructWithDb) SelectViaQueryx() error {
	rows, err := e.db.Queryx("SELECT * FROM t1")
	if err != nil {
		return err
	}

	var row Row
	for rows.Next() {
		if err := rows.StructScan(&row); err != nil {
			return err
		}

		log.Printf("%+v", row)
	}

	// close rows in defer
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// or just return result of
	return rows.Close()
}

func (e *ExampleStructWithDb) SelectViaScan() error {
	rows, err := e.db.Queryx("SELECT * FROM t1")
	if err != nil {
		return err
	}

	var id uint64
	var placebo string
	for rows.Next() {
		if err := rows.Scan(&id, &placebo); err != nil {
			return err
		}

		log.Printf("%d - %s", id, placebo)
	}

	// close rows in defer
	defer func() {
		if err := rows.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// or just return result of
	return rows.Close()
}

func (e *ExampleStructWithDb) SelectViaSelect() error {
	var rows []*Row
	// inside Select method is  Queryx
	if err := e.db.Select(&rows, "SELECT * FROM t1"); err != nil {
		return err
	}

	for _, row := range rows {
		log.Printf("%+v", row)
	}

	return nil
}

package main

import (
	"database/sql"
	"log"
	"testing"
	"time"

	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/jmoiron/sqlx"
)

type Row struct {
	Id        uint64    `db:"Id"`
	EventDate time.Time `db:"EventDate"`
	Data0     string    `db:"Data0"`
	Data1     string    `db:"Data1"`
	Data2     string    `db:"Data2"`
	Data3     string    `db:"Data3"`
	Data4     string    `db:"Data4"`
	Data5     string    `db:"Data5"`
	Data6     string    `db:"Data6"`
	Data7     string    `db:"Data7"`
	Data8     string    `db:"Data8"`
	Data9     string    `db:"Data9"`
	Data10    string    `db:"Data10"`
	Data11    string    `db:"Data11"`
	Data12    string    `db:"Data12"`
	Data13    string    `db:"Data13"`
	Data14    string    `db:"Data14"`
	Data15    string    `db:"Data15"`
	Data16    string    `db:"Data16"`
	Data17    string    `db:"Data17"`
	Data18    string    `db:"Data18"`
	Data19    string    `db:"Data19"`
	Data20    string    `db:"Data20"`
	Data21    string    `db:"Data21"`
	Data22    string    `db:"Data22"`
	Data23    string    `db:"Data23"`
	Data24    string    `db:"Data24"`
	Data25    string    `db:"Data25"`
	Data26    string    `db:"Data26"`
	Data27    string    `db:"Data27"`
	Data28    string    `db:"Data28"`
	Data29    string    `db:"Data29"`
}

func TestSelect(t *testing.T) {
	c, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?debug=false")
	if err != nil {
		log.Fatal(err)
	}

	var all []Row

	q, err := c.Query(`SELECT * FROM events`)
	if err != nil {
		t.Fatal(err)
	}

	for q.Next() {
		var row Row
		if err := q.Scan(&row.Id, &row.EventDate,
			&row.Data0,
			&row.Data1,
			&row.Data2,
			&row.Data3,
			&row.Data4,
			&row.Data5,
			&row.Data6,
			&row.Data7,
			&row.Data8,
			&row.Data9,
			&row.Data10,
			&row.Data11,
			&row.Data12,
			&row.Data13,
			&row.Data14,
			&row.Data15,
			&row.Data16,
			&row.Data17,
			&row.Data18,
			&row.Data19,
			&row.Data20,
			&row.Data21,
			&row.Data22,
			&row.Data23,
			&row.Data24,
			&row.Data25,
			&row.Data26,
			&row.Data27,
			&row.Data28,
			&row.Data29,
		); err != nil {
			t.Fatal(err)
		}
		all = append(all, row)
	}

	t.Log(all[0])
}

func BenchmarkSelect(b *testing.B) {
	c, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?debug=false")
	if err != nil {
		b.Fatal(err)
	}
	if err := c.Ping(); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		var all []Row

		q, err := c.Query(`SELECT * FROM events`)
		if err != nil {
			b.Fatal(err)
		}

		for q.Next() {
			var row Row
			if err := q.Scan(&row.Id, &row.EventDate,
				&row.Data0,
				&row.Data1,
				&row.Data2,
				&row.Data3,
				&row.Data4,
				&row.Data5,
				&row.Data6,
				&row.Data7,
				&row.Data8,
				&row.Data9,
				&row.Data10,
				&row.Data11,
				&row.Data12,
				&row.Data13,
				&row.Data14,
				&row.Data15,
				&row.Data16,
				&row.Data17,
				&row.Data18,
				&row.Data19,
				&row.Data20,
				&row.Data21,
				&row.Data22,
				&row.Data23,
				&row.Data24,
				&row.Data25,
				&row.Data26,
				&row.Data27,
				&row.Data28,
				&row.Data29,
			); err != nil {
				b.Fatal(err)
			}
			all = append(all, row)
		}
	}
}

func BenchmarkSelectx(b *testing.B) {
	c, err := sqlx.Open("clickhouse", "tcp://127.0.0.1:9000?debug=false")
	if err != nil {
		b.Fatal(err)
	}
	if err := c.Ping(); err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		var all []Row

		q, err := c.Queryx(`SELECT * FROM events`)
		if err != nil {
			b.Fatal(err)
		}

		for q.Next() {
			var row Row
			if err := q.StructScan(&row); err != nil {
				b.Fatal(err)
			}
			all = append(all, row)
		}
	}
}

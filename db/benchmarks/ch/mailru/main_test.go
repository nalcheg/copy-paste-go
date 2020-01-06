package main

import (
	"log"
	"testing"
	"time"

	"github.com/mailru/dbr"
	_ "github.com/mailru/go-clickhouse"
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
	c, err := dbr.Open("clickhouse", "http://127.0.0.1:8123/?debug=0", nil)
	if err != nil {
		log.Fatal(err)
	}
	sess := c.NewSession(nil)

	var rows []Row

	if err := sess.SelectBySql(`SELECT * FROM events LIMIT 514`).LoadValue(&rows); err != nil {
		t.Error(err)
	}

	log.Print(rows)
}

func BenchmarkSelect(b *testing.B) {
	c, err := dbr.Open("clickhouse", "http://127.0.0.1:8123/?debug=0", nil)
	if err != nil {
		log.Fatal(err)
	}
	sess := c.NewSession(nil)

	for i := 0; i < b.N; i++ {
		var rows []Row
		if err := sess.SelectBySql(`SELECT * FROM events`).LoadValue(&rows); err != nil {
			b.Error(err)
		}
	}
}

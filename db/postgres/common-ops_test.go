package main

import (
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func Benchmark_SelectViaSelect(b *testing.B) {
	b.ReportAllocs()
	dbDSN := "user=postgres password=postgres host=127.0.0.1 port=55432 dbname=test sslmode=disable"

	dbx, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		b.Fatal(err)
	}

	s := &ExampleStructWithDb{db: dbx}

	for i := 0; i < b.N; i++ {
		if err := s.SelectViaSelect(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_SelectViaScan(b *testing.B) {
	b.ReportAllocs()
	dbDSN := "user=postgres password=postgres host=127.0.0.1 port=55432 dbname=test sslmode=disable"

	dbx, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		b.Fatal(err)
	}

	s := &ExampleStructWithDb{db: dbx}

	for i := 0; i < b.N; i++ {
		if err := s.SelectViaScan(); err != nil {
			b.Fatal(err)
		}
	}
}

func Benchmark_SelectViaQueryx(b *testing.B) {
	b.ReportAllocs()
	dbDSN := "user=postgres password=postgres host=127.0.0.1 port=55432 dbname=test sslmode=disable"

	dbx, err := sqlx.Open("pgx", dbDSN)
	if err != nil {
		b.Fatal(err)
	}

	s := &ExampleStructWithDb{db: dbx}

	for i := 0; i < b.N; i++ {
		if err := s.SelectViaQueryx(); err != nil {
			b.Fatal(err)
		}
	}
}

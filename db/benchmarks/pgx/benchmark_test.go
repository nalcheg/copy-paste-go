package main

import (
	"context"
	"log"
	"testing"
)

func BenchmarkPgx(b *testing.B) {
	pgxPool, err := connectPgx()
	if err != nil {
		log.Fatal(err)
	}
	c, err := pgxPool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		if _, err := selectTenCommentsPgx(c); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkSqlx(b *testing.B) {
	dbx, err := connectSqlxWithPgx()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		if _, err := selectTenComments(dbx); err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkGorm(b *testing.B) {
	gormConn, err := connectGorm()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		selectTenCommentsGorm(gormConn)
	}
}

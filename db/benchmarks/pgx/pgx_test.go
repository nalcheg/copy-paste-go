package main

import (
	"context"
	"log"
	"testing"
)

func TestPgx(t *testing.T) {
	pgxPool, err := connectPgx()
	if err != nil {
		log.Fatal(err)
	}

	c, err := pgxPool.Acquire(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	comments, err := selectTenCommentsPgx(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", comments)
}

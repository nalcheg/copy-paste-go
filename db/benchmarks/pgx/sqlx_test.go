package main

import (
	"log"
	"testing"
)

func Test_selectTenComments(t *testing.T) {
	dbx, err := connectSqlxWithPgx()
	if err != nil {
		return
	}

	if _, err := selectTenComments(dbx); err != nil {
		log.Fatal(err)
	}
}

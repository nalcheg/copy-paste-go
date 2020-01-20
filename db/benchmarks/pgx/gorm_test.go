package main

import (
	"log"
	"testing"
)

func TestTest(t *testing.T) {
	gormConn, err := connectGorm()
	if err != nil {
		log.Fatal(err)
	}

	comments := selectTenCommentsGorm(gormConn)

	log.Printf("%+v", comments)
}

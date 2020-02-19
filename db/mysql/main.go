package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// DB_DSN=user:password@tcp(127.0.0.1)/database
func main() {
	db, err := sql.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}

	rows, err := db.Query("SELECT COUNT(id) FROM table_name")
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()

	var c int
	if err := rows.Scan(&c); err != nil {
		log.Fatal(err)
	}

	log.Printf("there are %d record in table_name", c)

	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}
}

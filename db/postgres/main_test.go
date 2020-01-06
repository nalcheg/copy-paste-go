package main

import (
	"context"
	"testing"
)

func TestMigrateWithDropDatabase(t *testing.T) {
	dbDSN := "user=postgres password=postgres host=127.0.0.1 port=55432 dbname=test sslmode=disable"

	db, err := databaseConnect(dbDSN)
	if err != nil {
		t.Error(err)
	}
	defer db.Close()

	if err := databaseMigrate(db); err != nil {
		t.Error(err)
	}

	if _, err := db.Exec(
		context.Background(),
		`INSERT INTO t1 (placebo) VALUES ($1), ($2), ($3)`,
		"placebo",
		"placeTwo",
		"placeThree",
	); err != nil {
		t.Error(err)
	}

	//if _, err := db.Exec(
	//	context.Background(),
	//	`DROP TABLE t1, t2, schema_version`,
	//); err != nil {
	//	t.Error(err)
	//}
	//
	//dbDSN = "user=postgres password=postgres host=127.0.0.1 port=55432 sslmode=disable"
	//db, err = databaseConnect(dbDSN)
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//if _, err := db.Exec(
	//	context.Background(),
	//	`SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'test'`,
	//); err != nil {
	//	t.Error(err)
	//}
	//
	//if _, err := db.Exec(
	//	context.Background(),
	//	`DROP DATABASE test`,
	//); err != nil {
	//	t.Error(err)
	//}
}

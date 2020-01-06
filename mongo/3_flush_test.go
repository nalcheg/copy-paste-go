package main

import "testing"

func TestMongo_FlushCollection(t *testing.T) {
	m := Mongo{}
	if err := m.Connect("mongodb://localhost:27017"); err != nil {
		t.Error(err)
	}

	if err := m.FlushCollection(); err != nil {
		t.Error(err)
	}
}

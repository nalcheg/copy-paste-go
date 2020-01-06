package main

import (
	"os"
	"testing"
)

func TestMongo_LoadDump(t *testing.T) {
	m := Mongo{}
	if err := m.Connect("mongodb://localhost:27017"); err != nil {
		t.Error(err)
	}

	if err := m.LoadDump("data.json"); err != nil {
		t.Error(err)
	}
}

func TestMongo_UnloadDump(t *testing.T) {
	m := Mongo{}
	if err := m.Connect("mongodb://localhost:27017"); err != nil {
		t.Error(err)
	}

	filename := "tempdata.json"

	if _, err := os.Open(filename); err == nil {
		t.Error(os.ErrExist)
	}

	if err := m.UnloadDump(filename); err != nil {
		t.Error(err)
	}

	if _, err := os.Open(filename); err != nil {
		t.Error(os.ErrNotExist)
	}

	if err := os.Remove(filename); err != nil {
		t.Error(err)
	}

	if _, err := os.Open(filename); err == nil {
		t.Error(os.ErrExist)
	}
}

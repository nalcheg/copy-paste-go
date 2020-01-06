package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func Test_read(t *testing.T) {
	filename := "temp.txt"
	content := []byte("0123456789")
	if err := ioutil.WriteFile(filename, content, 0644); err != nil {
		t.Fatal(err)
	}

	read(filename)

	if err := os.Remove(filename); err != nil {
		t.Fatal(err)
	}
}

func Test_write(t *testing.T) {
	filename1 := "temp1.txt"
	filename2 := "temp2.txt"

	write(filename1, filename2)

	if err := os.Remove(filename1); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filename2); err != nil {
		t.Fatal(err)
	}
}

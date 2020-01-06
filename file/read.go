package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func read(filename string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	log.Print(string(dat))

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	if err != nil {
		panic(err)
	}

	log.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	o2, err := f.Seek(6, 0)
	if err != nil {
		panic(err)
	}

	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	if err != nil {
		panic(err)
	}

	log.Printf("%d bytes @ %d: ", n2, o2)
	log.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	if err != nil {
		panic(err)
	}

	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	if err != nil {
		panic(err)
	}

	log.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	if err != nil {
		panic(err)
	}

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	if err != nil {
		panic(err)
	}

	log.Printf("5 bytes: %s\n", string(b4))

	if err := f.Close(); err != nil {
		panic(err)
	}
}

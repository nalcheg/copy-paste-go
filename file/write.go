package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

func write(filename1, filename2 string) {
	d1 := []byte("0123456789")
	if err := ioutil.WriteFile(filename1, d1, 0644); err != nil {
		panic(err)
	}

	f, err := os.Create(filename2)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	if err != nil {
		panic(err)
	}

	log.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	log.Printf("wrote %d bytes\n", n3)

	if err := f.Sync(); err != nil {
		panic(err)
	}

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	log.Printf("wrote %d bytes\n", n4)

	if err := w.Flush(); err != nil {
		panic(err)
	}
}

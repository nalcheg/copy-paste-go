package concurrency

import (
	"log"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	var data int
	//var data *int

	for i := 0; i < 100; i++ {
		go func() {
			data++
			//*data++
		}()
	}

	time.Sleep(1000 * time.Millisecond)

	log.Print(data)
}

func TestGoroutineTwo(t *testing.T) {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		//time.Sleep(100 * time.Millisecond)
		log.Printf("the value is %v.\n", data)
	}
}

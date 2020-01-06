package main

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"
)

func Test(t *testing.T) {
	serverAddres := "http://127.0.0.1:58080"
	app := App{cc: make(chan struct{})}

	go func() {
		time.Sleep(2 * time.Second)

		client := http.Client{}

		log.Print("start request 1")
		resp, err := client.Get(serverAddres)
		if err != nil {
			log.Print(err)
		}
		if resp.StatusCode != 204 {
			t.Fatal("not expected response code 1")
		}
		log.Print("request 1 respose code: ", resp.StatusCode)

		func() {
			log.Print("start request 2")
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			code, err := getWithContext(ctx, serverAddres)
			if err != nil {
				log.Print(err)
			}
			if code != 204 {
				t.Fatal("not expected response code 2")
			}
			log.Print("request 2 respose code: ", code)
		}()

		func() {
			log.Print("start request 3")
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			code, err := getWithContext(ctx, serverAddres+"/sleep")
			if err != nil {
				log.Print(err)
			}
			if code != 408 {
				t.Fatal("not expected response code 3")
			}
			log.Print("request 3 respose code: ", code)
		}()
	}()

	go func() {
		time.Sleep(8 * time.Second)
		app.cc <- struct{}{}
	}()

	app.Start(":58080")
}

func getWithContext(ctx context.Context, url string) (int, error) {
	resultChannel := make(chan int)

	go func() {
		client := http.Client{}
		resp, err := client.Get(url)
		if err != nil {
			ctx.Done()
			return
		}
		resultChannel <- resp.StatusCode
	}()

	select {
	case <-ctx.Done():
		return 408, ctx.Err()
	case code := <-resultChannel:
		return code, nil
	}
}

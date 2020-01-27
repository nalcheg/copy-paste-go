package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	app := App{
		cc: make(chan struct{}),
	}
	app.Start(":58080")
}

type App struct {
	cc chan struct{}
}

func (a *App) Start(addr string) {
	gofakeit.Seed(0)

	r := router.New()
	r.GET("/", now)
	r.GET("/sleep", sleep)
	r.GET("/beer", fastHTTPHandlerWithTimeout)

	go func() {
		log.Fatal(fasthttp.ListenAndServe(addr, r.Handler))
	}()

	for {
		select {
		case _ = <-a.cc:
			log.Print("exiting")
			//log.New()
			return
		}
	}
}

func now(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(204)
}

func sleep(ctx *fasthttp.RequestCtx) {
	time.Sleep(5 * time.Second)
	ctx.SetStatusCode(204)
}

type ResultData struct {
	Data string `json:"data"`
}

func fastHTTPHandlerWithTimeout(ctx *fasthttp.RequestCtx) {
	c, cancel := context.WithTimeout(ctx, 100*time.Millisecond)
	defer cancel()

	resp := make(chan ResultData)

	go process(c, resp)

	var res ResultData
	select {
	case <-c.Done():
		fmt.Printf("error: processing timeout \n")
		break
	case res = <-resp:
		fmt.Printf("finished processing\n")
	}

	if res.Data == "" {
		ctx.SetStatusCode(http.StatusRequestTimeout)
		return
	}

	r, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	if _, err := ctx.Write(r); err != nil {
		panic(err)
	}
	ctx.SetStatusCode(http.StatusOK)
}

func process(ctx context.Context, respChan chan<- ResultData) {
	randomTimeout := gofakeit.Number(50, 150)

	time.Sleep(time.Duration(randomTimeout) * time.Millisecond)

	respChan <- ResultData{Data: gofakeit.BeerName()}
}

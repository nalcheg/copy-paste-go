package main

import (
	"log"
	"time"

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
	r := router.New()
	r.GET("/", now)
	r.GET("/sleep", sleep)

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

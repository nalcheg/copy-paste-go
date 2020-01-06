package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/nalcheg/copy-paste-go/http/fasthttp/middleware"
)

func index(ctx *fasthttp.RequestCtx) {
	if _, err := fmt.Fprint(ctx, "Who are you?"); err != nil {
		log.Printf("%v", err)
	}
}

func welcome(ctx *fasthttp.RequestCtx) {
	if _, err := fmt.Fprintf(ctx, "Hi, %s !", ctx.FormValue("iam")); err != nil {
		log.Printf("%v", err)
	}
}

func main() {
	r := router.New()
	r.GET("/", index)
	r.GET("/welcome", welcome)

	srv := &fasthttp.Server{
		Handler: middleware.Get(r.Handler),
	}

	go func() {
		log.Fatal(srv.ListenAndServe(":9011"))
	}()

	waitForShutdown(srv)
}

func waitForShutdown(srv *fasthttp.Server) {
	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interruptChan

	if err := srv.Shutdown(); err != nil {
		log.Fatal(err)
	}

	log.Println("Shutting down")
	os.Exit(0)
}

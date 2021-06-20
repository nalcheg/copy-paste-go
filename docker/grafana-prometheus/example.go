package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("%#v", http.ListenAndServe(":80", nil))
}

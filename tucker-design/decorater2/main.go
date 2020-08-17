package main

import (
	"log"
	"net/http"
	"time"

	"github.com/tuckersGo/goWeb/web10/decoHandler"
	"github.com/tuckersGo/goWeb/web10/myapp"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time: ", time.Since(start).Nanoseconds())
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	h := decoHandler.NewDecoHandler(mux, logger)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}

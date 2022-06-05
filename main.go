package main

import (
	"flag"
	"net/http"

	_ "embed"

	"log"

	"goji.io"
	"goji.io/pat"
)

var address string

//go:embed body.html
var body []byte

func redirectWithMeta(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Location", "https://github.com/hop-proto/hop-go")
	w.WriteHeader(302)
	_, _ = w.Write(body)
}

func main() {
	flag.StringVar(&address, "address", ":8080", "listen address")
	flag.Parse()

	mux := goji.NewMux()
	mux.Handle(pat.Get("/hop"), http.HandlerFunc(redirectWithMeta))

	log.Printf("listening on %s", address)
	_ = http.ListenAndServe(address, mux)
}

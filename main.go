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

func redirectToWithMeta(destination string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", destination)
		w.WriteHeader(302)
		_, _ = w.Write(body)
	}
}

func main() {
	flag.StringVar(&address, "address", ":8080", "listen address")
	flag.Parse()

	mux := goji.NewMux()
	mux.Handle(pat.Get("/hop"), redirectToWithMeta("https://github.com/hop-proto/hop-go"))
	mux.Handle(pat.Get("/vend"), redirectToWithMeta("https://github.com/hop-proto/hop-vend"))

	log.Printf("listening on %s", address)
	_ = http.ListenAndServe(address, mux)
}

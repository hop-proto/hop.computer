package main

import (
	"bytes"
	"flag"
	"html/template"
	"net/http"

	_ "embed"

	"log"

	"goji.io"
	"goji.io/pat"
)

var address string

//go:embed body.html
var bodyHTML string

var bodyTemplate *template.Template

func init() {
	tmpl, err := template.New("body").Parse(string(bodyHTML))
	if err != nil {
		log.Fatalf("failed to parse body.html template: %v", err)
	}
	bodyTemplate = tmpl
}

func redirectToWithMeta(destination string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", destination)
		w.WriteHeader(302)

		var buf bytes.Buffer
		err := bodyTemplate.Execute(&buf, map[string]string{
			"PackageName": "hop.computer" + r.URL.Path,
			"GitURL":      destination,
		})
		if err != nil {
			http.Error(w, "template execution error", http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(buf.Bytes())
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

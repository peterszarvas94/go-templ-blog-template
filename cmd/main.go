package main

import (
	"net/http"
	"peterszarvas94/blog/templates"

	"github.com/a-h/templ"
)

func main() {
	indexPage := templates.Index("Peter")

	http.Handle("/", templ.Handler(indexPage))
	http.ListenAndServe("localhost:8080", nil)
}

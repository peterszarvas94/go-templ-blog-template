package main

import (
	"net/http"
	"peterszarvas94/blog/pages"
	"peterszarvas94/blog/pkg"

	"github.com/a-h/templ"
)

type indexHandler struct{}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	files := pkg.GetFiles()
	handler := templ.Handler(pages.Index(files))
	handler.ServeHTTP(w, r)
	return
}

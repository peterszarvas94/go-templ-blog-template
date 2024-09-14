package main

import (
	"net/http"
	"peterszarvas94/blog/pkg/fileutils"
	"peterszarvas94/blog/pkg/pages"

	"github.com/a-h/templ"
)

type indexHandler struct{}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	files := fileutils.GetFiles()
	handler := templ.Handler(pages.Index(files))
	handler.ServeHTTP(w, r)
	return
}

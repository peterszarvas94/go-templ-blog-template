package main

import (
	"net/http"
	"peterszarvas94/blog/pkg/custom"
	"peterszarvas94/blog/pkg/fileutils"
	"peterszarvas94/blog/pkg/pages"

	"github.com/a-h/templ"
)

type indexHandler struct{}

func (h *indexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if customHomePage, exists := (*custom.Routes)["/"]; exists {
		handler := templ.Handler(customHomePage)
		handler.ServeHTTP(w, r)
	} else {
		files := fileutils.GetFiles()
		handler := templ.Handler(pages.Index(files))
		handler.ServeHTTP(w, r)
	}
	return
}

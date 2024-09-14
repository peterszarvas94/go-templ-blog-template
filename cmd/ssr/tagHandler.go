package main

import (
	"net/http"
	"peterszarvas94/blog/pkg/fileutils"
	"peterszarvas94/blog/pkg/pages"

	"github.com/a-h/templ"
)

type tagHandler struct{}

func (h *tagHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tag := r.PathValue("tag")

	tags := fileutils.GetTags()
	files := tags[tag]

	handler := templ.Handler(pages.NotFound())

	if len(files) > 0 {
		handler = templ.Handler(pages.Tag(tag, files))
	}

	handler.ServeHTTP(w, r)
	return
}

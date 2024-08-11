package main

import (
	"net/http"
	"peterszarvas94/blog/pkg"
	"peterszarvas94/blog/templates"

	"github.com/a-h/templ"
)

func handleTag(w http.ResponseWriter, r *http.Request) {
	tag := r.PathValue("tag")
	if tag == "" {
		http.Error(w, "missing tag", http.StatusBadRequest)
		return
	}

	// files := pkg.FindFilesByTag(tag)
	tags := pkg.CollectTags()
	files := tags[tag]

	tagPage := templates.TagPage(tag, files)
	handler := templ.Handler(tagPage)

	handler.ServeHTTP(w, r)
}

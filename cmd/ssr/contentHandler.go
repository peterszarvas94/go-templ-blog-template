package main

import (
	"fmt"
	"net/http"
	"path"
	"peterszarvas94/blog/config"
	"peterszarvas94/blog/pages"
	"peterszarvas94/blog/pkg"

	"github.com/a-h/templ"
)

type contentHandler struct{}

func (h *contentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: assert CheckContentDir has run

	// TODO: only for dev
	r.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate")
	r.Header.Set("Pragma", "no-cache")
	r.Header.Set("Expires", "0")

	url := r.URL.Path

	pathToFile := path.Join(config.Dirs.Content, url)

	pathToFileWithExtension := fmt.Sprintf("%s.md", pathToFile)

	file, err := pkg.FindFileFromFilePath(pathToFileWithExtension)
	if err != nil {
		handler := templ.Handler(pages.NotFound())
		handler.ServeHTTP(w, r)
		return
	}

	handler := templ.Handler(pages.Post(file))
	handler.ServeHTTP(w, r)
}

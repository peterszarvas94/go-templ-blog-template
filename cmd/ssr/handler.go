package main

import (
	"fmt"
	"net/http"
	"path"
	"peterszarvas94/blog/config"
	"peterszarvas94/blog/pages"
	"peterszarvas94/blog/pkg"
	"strings"

	"github.com/a-h/templ"
)

type contentHandler struct{}

func (h *contentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: assert CheckContentDir has run
	url := r.URL.Path

	if url == "/" {
		files := pkg.GetFiles()
		handler := templ.Handler(pages.Index(files))
		handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/tag") {
		tag := path.Base(url)
		fmt.Println("tag", tag)

		tags := pkg.GetTags()
		files := tags[tag]
		handler := templ.Handler(pages.Tag(tag, files))
		handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/category") {
		// TODO: category
	}

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

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

	// TODO: only for dev
	r.Header.Set("Cache-Control", "no-store, no-cache, must-revalidate")
	r.Header.Set("Pragma", "no-cache")
	r.Header.Set("Expires", "0")

	url := r.URL.Path

	if url == "/" {
		files := pkg.GetFiles()
		handler := templ.Handler(pages.Index(files))
		handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/tag") {
		tag := path.Base(url)

		tags := pkg.GetTags()
		files := tags[tag]
		fmt.Println(files)

		handler := templ.Handler(pages.NotFound())

		if len(files) > 0 {
			handler = templ.Handler(pages.Tag(tag, files))
		}

		handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/category") {
		category := path.Base(url)

		categories := pkg.GetCategories()
		files := categories[category]

		handler := templ.Handler(pages.Category(category, files))
		handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/search") {
		handler := templ.Handler(pages.Search())
		handler.ServeHTTP(w, r)
		return
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

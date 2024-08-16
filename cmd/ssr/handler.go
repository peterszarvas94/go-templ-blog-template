package main

import (
	"fmt"
	"net/http"
	"path"
	"peterszarvas94/blog/pkg"
	"peterszarvas94/blog/templates"
	"strings"

	"github.com/a-h/templ"
)

type ContentHandler struct{}

type Pages struct {
	notFoundPage pkg.NotFoundPage
	indexPage    pkg.IndexPage
	postPage     pkg.PostPage
	tagPage      pkg.TagPage
	categoryPage pkg.CategoryPage
}

var components = &Pages{
	notFoundPage: templates.NotFoundPage,
	indexPage:    templates.IndexPage,
	postPage:     templates.PostPage,
	tagPage:      templates.TagPage,
	categoryPage: templates.CategoryPage,
}

func (h *ContentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: assert CheckContentDir has run
	url := r.URL.Path

	if url == "/" {
		files := pkg.GetFiles()
		handler := templ.Handler(components.indexPage(files))
		handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/tag") {
		tag := path.Base(url)
		fmt.Println("tag", tag)

		tags := pkg.CollectTags()
		files := tags[tag]
		handler := templ.Handler(components.tagPage(tag, files))
		handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/category") {
		// TODO: category
	}

	pathToFile := path.Join(pkg.Config.ContentDir, url)

	pathToFileWithExtension := fmt.Sprintf("%s.md", pathToFile)

	file, err := pkg.FindFileFromFilePath(pathToFileWithExtension)
	if err != nil {
		handler := templ.Handler(components.notFoundPage())
		handler.ServeHTTP(w, r)
		return
	}

	handler := templ.Handler(components.postPage(file))
	handler.ServeHTTP(w, r)
}

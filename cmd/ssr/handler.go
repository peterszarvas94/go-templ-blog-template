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

type fileExists struct {
	exists bool
	file   *pkg.FileData
}

type ContentHandler struct{}

func (h *ContentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: assert CheckContentDir has run
	url := r.URL.Path

	if url == "/" {
		files := pkg.GetFiles()
		indexPage := templates.IndexPage(files)
		indexHandler := templ.Handler(indexPage)
		indexHandler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/tag") {
		// TODO: handle tag
		tag := strings.TrimPrefix(url, "/tag/")
		fmt.Println("tag", tag)

		tags := pkg.CollectTags()
		files := tags[tag]

		tagPage := templates.TagPage(tag, files)
		handler := templ.Handler(tagPage)

		handler.ServeHTTP(w, r)
		return
	}

	if strings.HasPrefix(url, "/category") {
		// TODO: handle tag
	}

	pathToFile := path.Join(pkg.Config.ContentDir, url)

	pathToFileWithExtension := fmt.Sprintf("%s.md", pathToFile)

	file, err := pkg.FindFileFromFilePath(pathToFileWithExtension)
	if err != nil {
		notFoundPage := templates.NotFoundPage()
		handler := templ.Handler(notFoundPage)

		handler.ServeHTTP(w, r)
		return
	}

	postPage := templates.PostPage(file)
	handler := templ.Handler(postPage)

	handler.ServeHTTP(w, r)
}

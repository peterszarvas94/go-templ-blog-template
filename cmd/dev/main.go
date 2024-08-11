package main

import (
	"fmt"
	"net/http"
	"peterszarvas94/blog/pkg"
	"peterszarvas94/blog/templates"

	"github.com/a-h/templ"
)

func main() {
	err := pkg.LoadEnvs()
	if err != nil {
		fmt.Println(err)
	}

	files, err := pkg.CollectFiles()
	if err != nil {
		fmt.Println(err)
	}

	fs := http.FileServer(http.Dir(pkg.Config.STATIC_DIR))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	notFoundPage := templates.NotFoundPage()
	http.Handle("/", templ.Handler(notFoundPage))

	indexPage := templates.IndexPage(files)
	http.Handle("/{$}", templ.Handler(indexPage))

	http.Handle("/posts/{year}/{month}/{day}/{filename}/{$}", http.HandlerFunc(handlePost))

	http.Handle("/tags/{tag}/{$}", http.HandlerFunc(handleTag))

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

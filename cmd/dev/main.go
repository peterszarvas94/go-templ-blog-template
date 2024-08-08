package main

import (
	"fmt"
	"net/http"
	"peterszarvas94/blog/pkg"
	"peterszarvas94/blog/templates"

	"github.com/a-h/templ"
)

func main() {
	files, err := pkg.CollectFiles()
	if err != nil {
		fmt.Println(err)
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	notFoundPage := templates.NotFoundPage()
	http.Handle("/", templ.Handler(notFoundPage))

	indexPage := templates.IndexPage(files)
	http.Handle("/{$}", templ.Handler(indexPage))

	http.Handle("/posts/{year}/{month}/{day}/{filename}/{$}", http.HandlerFunc(handlePost))

	http.ListenAndServe("localhost:8080", nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	year := r.PathValue("year")
	if year == "" {
		http.Error(w, "missing year", http.StatusBadRequest)
		return
	}

	month := r.PathValue("month")
	if month == "" {
		http.Error(w, "missing month", http.StatusBadRequest)
		return
	}

	day := r.PathValue("day")
	if day == "" {
		http.Error(w, "missing day", http.StatusBadRequest)
		return
	}

	filename := r.PathValue("filename")
	if filename == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}

	file, err := pkg.FindFile(year, month, day, filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	postPage := templates.PostPage(file)
	handler := templ.Handler(postPage)

	handler.ServeHTTP(w, r)
}

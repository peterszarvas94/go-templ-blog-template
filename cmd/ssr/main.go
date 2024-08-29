package main

import (
	"fmt"
	"net/http"
	"path"
	_ "peterszarvas94/blog/pkg/init"
)

func main() {
	staticDir := path.Join("theme", "static")

	fs := http.FileServer(http.Dir(staticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	indexHandler := &indexHandler{}
	http.Handle("/{$}", indexHandler)

	categoryHandler := &categoryHandler{}
	http.Handle("/category/{category}/{$}", categoryHandler)

	tagHandler := &tagHandler{}
	http.Handle("/tag/{tag}/{$}", tagHandler)

	contentHandler := &contentHandler{}
	http.Handle("/{segments...}", contentHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

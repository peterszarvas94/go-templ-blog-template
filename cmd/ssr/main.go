package main

import (
	"fmt"
	"net/http"
	"peterszarvas94/blog/config"
	_ "peterszarvas94/blog/pkg"
)

func main() {
	fs := http.FileServer(http.Dir(config.Dirs.Static))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	contentHandler := &contentHandler{}
	http.Handle("/{segments...}", contentHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

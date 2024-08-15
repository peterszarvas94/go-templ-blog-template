package main

import (
	"fmt"
	"net/http"
	"peterszarvas94/blog/pkg"
)

func main() {
	err := pkg.CheckContentDir()
	if err != nil {
		panic(err)
	}

	err = pkg.CollectFiles()
	if err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir(pkg.Config.StaticDir))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "favicon.ico")
	})

	contentHandler := &ContentHandler{}
	http.Handle("/{segments...}", contentHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

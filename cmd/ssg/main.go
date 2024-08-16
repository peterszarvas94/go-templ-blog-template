package main

import (
	"fmt"
	"net/http"
	"peterszarvas94/blog/config"
)

func main() {
	fs := http.FileServer(http.Dir(config.Dirs.Public))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

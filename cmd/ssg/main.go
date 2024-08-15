package main

import (
	"fmt"
	"net/http"
	"peterszarvas94/blog/pkg"
)

func main() {
	fs := http.FileServer(http.Dir(pkg.Config.PublicDir))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

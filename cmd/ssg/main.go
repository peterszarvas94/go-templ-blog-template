package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"peterszarvas94/blog/config"
)

func main() {
	fs := http.FileServer(http.Dir(config.Dirs.Public))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filePath := filepath.Join(config.Dirs.Public, r.URL.Path)

		fileInfo, err := os.Stat(filePath)
		if os.IsNotExist(err) || (err == nil && fileInfo.IsDir()) {
			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, filepath.Join(config.Dirs.Public, "404.html"))
			return
		}

		fs.ServeHTTP(w, r)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

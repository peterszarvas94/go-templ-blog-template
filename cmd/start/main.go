package main

import (
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	// Define the base directory to serve files from
	baseDir := "./public" // Assuming your static files are in the 'public' directory

	// Create a file server handler for serving static files
	fileServer := http.FileServer(http.Dir(baseDir))

	// Create a custom handler to map "/dir" -> "/dir/index.html"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Check if it's a directory and serve the corresponding index.html
		fullPath := filepath.Join(baseDir, path)
		if stat, err := os.Stat(fullPath); err == nil && stat.IsDir() {
			// Ensure the path ends with a slash
			if path[len(path)-1] != '/' {
				http.Redirect(w, r, path+"/", http.StatusMovedPermanently)
				return
			}

			// Serve the index.html file
			fullPath = filepath.Join(fullPath, "index.html")
			if _, err := os.Stat(fullPath); err == nil {
				http.ServeFile(w, r, fullPath)
				return
			}
		}

		// If not a directory or no index.html, fall back to the file server
		fileServer.ServeHTTP(w, r)
	})

	// Start the server
	http.ListenAndServe("localhost:8080", nil)
}


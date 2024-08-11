package main

import (
	"fmt"
	"net/http"
	"peterszarvas94/blog/pkg"
)

func main() {
	err := pkg.LoadEnvs()
	if err != nil {
		fmt.Println(err)
	}

	fs := http.FileServer(http.Dir(pkg.Config.PUBLIC_DIR))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe("localhost:8080", nil)
}

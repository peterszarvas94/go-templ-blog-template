package main

import (
	// "bytes"
	// "log"
	"net/http"
	// "path"
	// "peterszarvas94/blog/pkg"
	//
	// "github.com/gosimple/slug"
	// "github.com/yuin/goldmark"
)

func main() {

	rootPath := "public"

	fs := http.FileServer(http.Dir(rootPath))

	// for _, post := range pkg.Posts {
	// 	dir := path.Join(rootPath, post.Date.Format("2006/01/02"), slug.Make(post.Title))
	// 	// if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
	// 	// 	log.Fatalf("failed to create dir %q: %v", dir, err)
	// 	// }
	//
	// 	// name := path.Join(dir, "index.html")
	// 	// f, err := os.Create(name)
	// 	// if err != nil {
	// 	// 	log.Fatalf("failed to create output file: %v", err)
	// 	// }
	//
	// 	var buf bytes.Buffer
	// 	if err := goldmark.Convert([]byte(post.Content), &buf); err != nil {
	// 		log.Fatalf("failed to convert markdown to HTML: %v", err)
	// 	}
	//
	// 	content := pkg.Unsafe(buf.String())
	//
	// 	err := templates.ContentPage(post.Title, content).Render(context.Background(), f)
	// 	if err != nil {
	// 		log.Fatalf("failed to write output file: %v", err)
	// 	}
	// }

	http.Handle("/", fs)

	http.ListenAndServe("localhost:8080", nil)
}

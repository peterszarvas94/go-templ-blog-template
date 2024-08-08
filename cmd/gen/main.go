package main

import (
	"context"
	"fmt"
	"os"
	"path"

	// "net/http"
	"peterszarvas94/blog/pkg"
	"peterszarvas94/blog/templates"
	//
	// "github.com/a-h/templ"
)

func main() {
	files, err := pkg.CollectFiles()
	if err != nil {
		fmt.Println(err)
	}

	output := "public"
	if _, err := os.Stat(output); !os.IsNotExist(err) {
		err = os.RemoveAll(output)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Removed existing directory: ", output)
	}

	err = os.Mkdir(output, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Generated public directory: ", output)

	indexFileName := path.Join(output, "index.html")
	indexFile, err := os.Create(indexFileName)
	if err != nil {
		fmt.Println(err)
	}

	err = templates.IndexPage(files).Render(context.Background(), indexFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Generated index page:", indexFileName)

	for _, file := range *files {
		dir := path.Join(output, file.Date.Format("2006/01/02"), file.Filename)
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			fmt.Println(err)
		}

		postFileName := path.Join(dir, "index.html")
		postFile, err := os.Create(postFileName)
		if err != nil {
			fmt.Println(err)
		}

		err = templates.PostPage(&file).Render(context.Background(), postFile)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Generated post page:", postFileName)
	}

	fmt.Println("Done")
}

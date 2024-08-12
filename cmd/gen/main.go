package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"peterszarvas94/blog/pkg"
	"peterszarvas94/blog/templates"
)

func main() {
	output := pkg.Config.PublicDir

	if _, err := os.Stat(output); !os.IsNotExist(err) {
		err = os.RemoveAll(output)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Removed existing directory: ", output)
	}

	err := os.Mkdir(output, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Generated public directory: ", output)

	files, err := pkg.CollectFiles()
	if err != nil {
		fmt.Println(err)
	}

	// index
	indexFileName := path.Join(output, "index.html")
	indexFile, err := os.Create(indexFileName)
	if err != nil {
		fmt.Println(err)
	}

	err = templates.IndexPage(files).Render(context.Background(), indexFile)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Generated home page:", indexFileName)

	// posts
	for _, file := range *files {
		dir := path.Join(output, "posts", file.Date.Format("2006/01/02"), file.Filename)
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

	// tags
	tags := pkg.CollectTags()

	for tag, files := range tags {
		dir := path.Join(output, "tags", tag)
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			fmt.Println(err)
		}

		tagFileName := path.Join(dir, "index.html")
		tagFile, err := os.Create(tagFileName)
		if err != nil {
			fmt.Println(err)
		}

		err = templates.TagPage(tag, files).Render(context.Background(), tagFile)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Generated tag page:", tagFileName)
	}

	// static
	err = pkg.CopyFlatDir(pkg.Config.StaticDir, "public/static")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Copied static files:", "public/static")

	fmt.Println("Done")
}

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
	err := pkg.CheckContentDir()
	if err != nil {
		panic(err)
	}

	err = pkg.CollectFiles()
	if err != nil {
		panic(err)
	}

	publicDir := pkg.Config.PublicDir

	if _, err := os.Stat(publicDir); !os.IsNotExist(err) {
		err = os.RemoveAll(publicDir)
		if err != nil {
			panic(err)
		}
		fmt.Println("Removed existing directory: ", publicDir)
	}

	err = os.Mkdir(publicDir, 0755)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated public directory: ", publicDir)

	files := pkg.GetFiles()

	// index
	indexFileName := path.Join(publicDir, "index.html")
	indexFile, err := os.Create(indexFileName)
	if err != nil {
		panic(err)
	}

	err = templates.IndexPage(files).Render(context.Background(), indexFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated home page:", indexFileName)

	//... TODO continue

	// posts
	for _, file := range *files {
		dir := path.Join(publicDir, file.Filename)
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			panic(err)
		}

		postFileName := path.Join(dir, "index.html")
		postFile, err := os.Create(postFileName)
		if err != nil {
			panic(err)
		}

		err = templates.PostPage(&file).Render(context.Background(), postFile)
		if err != nil {
			panic(err)
		}

		fmt.Println("Generated post page:", postFileName)
	}

	// tags
	tags := pkg.CollectTags()

	for tag, files := range tags {
		dir := path.Join(publicDir, "tag", tag)
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			panic(err)
		}

		tagFileName := path.Join(dir, "index.html")
		tagFile, err := os.Create(tagFileName)
		if err != nil {
			panic(err)
		}

		err = templates.TagPage(tag, files).Render(context.Background(), tagFile)
		if err != nil {
			panic(err)
		}

		fmt.Println("Generated tag page:", tagFileName)
	}

	// static
	err = pkg.CopyFlatDir(pkg.Config.StaticDir, "public/static")
	if err != nil {
		panic(err)
	}

	fmt.Println("Copied static files:", "public/static")

	fmt.Println("Done")
}

package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"peterszarvas94/blog/config"
	"peterszarvas94/blog/pages"
	"peterszarvas94/blog/pkg"
)

func main() {
	publicDir := config.Dirs.Public

	if _, err := os.Stat(publicDir); !os.IsNotExist(err) {
		err = os.RemoveAll(publicDir)
		if err != nil {
			panic(err)
		}
		fmt.Println("Removed existing directory: ", publicDir)
	}

	err := os.Mkdir(publicDir, 0755)
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

	err = pages.Index(files).Render(context.Background(), indexFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated home page:", indexFileName)

	// posts
	for _, file := range files {
		dir := path.Join(publicDir, file.Fileroute)
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			panic(err)
		}

		postFileName := path.Join(dir, "index.html")
		postFile, err := os.Create(postFileName)
		if err != nil {
			panic(err)
		}

		err = pages.Post(file).Render(context.Background(), postFile)
		if err != nil {
			panic(err)
		}

		fmt.Println("Generated post page:", postFileName)
	}

	// tags
	tags := pkg.GetTags()

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

		err = pages.Tag(tag, files).Render(context.Background(), tagFile)
		if err != nil {
			panic(err)
		}

		fmt.Println("Generated tag page:", tagFileName)
	}

	// categories
	category := pkg.GetCategories()

	for category, files := range category {
		dir := path.Join(publicDir, "category", category)
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			panic(err)
		}

		categoryFileName := path.Join(dir, "index.html")
		categoryFile, err := os.Create(categoryFileName)
		if err != nil {
			panic(err)
		}

		err = pages.Category(category, files).Render(context.Background(), categoryFile)
		if err != nil {
			panic(err)
		}

		fmt.Println("Generated category page:", categoryFileName)
	}

	// search
	dir := path.Join(publicDir, "search")
	if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
		panic(err)
	}

	searchFileName := path.Join(dir, "index.html")
	searchFile, err := os.Create(searchFileName)
	if err != nil {
		panic(err)
	}

	err = pages.Search().Render(context.Background(), searchFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated search page:", searchFileName)

	// static
	err = CopyFlatDir(config.Dirs.Static, "public/static")
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
}

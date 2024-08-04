package main

import (
	"fmt"
	"peterszarvas94/blog/pkg"
)

func gen() {
	mdFiles, err := pkg.GetMdFiles()
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range *mdFiles {
		fmt.Println("--------")
		fmt.Println(file.Path)
		fmt.Printf("%+v\n", file.Matter)
		fmt.Println(file.Date.String())
		fmt.Println(file.Html)
	}
}

// in dev -> watch md and template files -> live server
// in prod -> generate htmls

package init

import (
	"fmt"
	"peterszarvas94/blog/pkg"
)

func init() {
	err := pkg.CheckContentDir()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Content directory is valid")

	err = pkg.CollectFiles()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Collected files")

	err = pkg.WriteFilesJsonFile()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Generated files.json")
}

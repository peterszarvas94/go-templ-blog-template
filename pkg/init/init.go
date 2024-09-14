package pkg

import (
	"fmt"
	"peterszarvas94/blog/pkg/check"
	"peterszarvas94/blog/pkg/config"
	"peterszarvas94/blog/pkg/fileutils"
)

func init() {
	err := check.CheckContentDir()
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Content directory is valid")

	if config.GeneretareFilesJson {
		err = fileutils.WriteFilesJsonFile()
		if err != nil {
			panic(err)
		}

		fmt.Println("✅ Generated files.json")
	}
}

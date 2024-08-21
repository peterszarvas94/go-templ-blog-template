package pkg

import "fmt"

func init() {
	err := CheckContentDir()
	if err != nil {
		panic(err)
	}

	fmt.Println("Content directory is valid")

	err = CollectFiles()
	if err != nil {
		panic(err)
	}

	fmt.Println("Collected files")

	err = WriteFilesJsonFile()
	if err != nil {
		panic(err)
	}

	fmt.Println("Generated files.json")
}

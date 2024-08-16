package pkg

import (
	"os"
)

var protectedDirs = []string{"static", "tag", "category"}

func CheckContentDir() error {
	folder, err := os.Open("content")
	if err != nil {
		return err
	}

	files, err := folder.Readdir(-1)
	if err != nil {
		return err
	}

	var dirs []string

	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}

	for _, dir := range dirs {
		for _, protectedDir := range protectedDirs {
			if dir == protectedDir {
				return &ProtectedDirError{dir}
			}
		}
	}

	return nil
}

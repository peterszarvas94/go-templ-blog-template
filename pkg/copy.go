package pkg

import (
	"io"
	"os"
	"path/filepath"
)

func copyFile(srcPath, dstPath string) error {
	inFile, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer inFile.Close()

	outFile, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, inFile)
	return err
}

func CopyFlatDir(srcDir, dstDir string) error {
	files, err := os.ReadDir(srcDir)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dstDir, os.ModePerm); err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			srcPath := filepath.Join(srcDir, file.Name())
			dstPath := filepath.Join(dstDir, file.Name())
			if err := copyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}
	return nil
}

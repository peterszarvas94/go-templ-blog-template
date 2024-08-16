package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type FileData struct {
	Filename string
	Matter   FrontMatter
	Date     time.Time
	Html     string
	Path     string
}

var files []*FileData

func FindFileFromFilePath(filePath string) (*FileData, error) {
	// TODO: assert collect files is run (not nil)

	for _, file := range files {
		if file.Path == filePath {
			return file, nil
		}
	}

	return nil, fmt.Errorf("file not found: %s", filePath)
}

func GetFiles() []*FileData {
	return files
}

func CollectFiles() error {
	files = files[:0]

	dirName := Config.ContentDir

	err := filepath.Walk(dirName, walkContentDir)
	if err != nil {
		return err
	}

	return nil
}

func walkContentDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
		return nil
	}

	contentBytes, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file: %s", path)
	}

	rawContent := string(contentBytes)
	matter, rawBody, err := parseFrontMatter(rawContent)
	if err != nil {
		return fmt.Errorf("failed to parse front matter: %s", path)
	}

	html, err := parseFileContent(rawBody)
	if err != nil {
		return fmt.Errorf("failed to parse file content: %s", path)
	}

	filename := strings.TrimSuffix(path, ".md")
	filename = strings.TrimPrefix(filename, Config.ContentDir)

	// Get the first part (before the dot)
	file := &FileData{
		Filename: filename,
		Matter:   matter,
		Date:     time.Now(), // TODO
		Html:     html,
		Path:     path,
	}
	files = append(files, file)

	return nil
}

package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type FrontMatter struct {
	Title    string   `yaml:"title"`
	Category string   `yaml:"category"`
	Tags     []string `yaml:"tags"`
}

type FileData struct {
	Filename string
	Matter   FrontMatter
	Date     time.Time
	Html     string
	Path     string
}

var files []FileData

func walkContentDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() || !strings.HasSuffix(info.Name(), ".md") {
		return nil
	}

	parts := strings.Split(path, string(os.PathSeparator))

	if len(parts) < 5 {
		return nil
	}

	year := parts[len(parts)-4]
	month := parts[len(parts)-3]
	day := parts[len(parts)-2]

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return fmt.Errorf("invalid year component in path: %s", path)
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return fmt.Errorf("invalid month component in path: %s", path)
	}

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		return fmt.Errorf("invalid day component in path: %s", path)
	}

	date := time.Date(yearInt, time.Month(monthInt), dayInt, 0, 0, 0, 0, time.UTC)

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

	fileName := info.Name()
	fileNameArr := strings.Split(fileName, ".")

	// Get the first part (before the dot)
	name := fileNameArr[0]

	file := FileData{
		Filename: name,
		Matter:   matter,
		Date:     date,
		Html:     html,
		Path:     path,
	}
	files = append(files, file)

	return nil
}

func CollectFiles() (*[]FileData, error) {
	files = files[:0]

	dirName := Config.CONTENT_DIR

	err := filepath.Walk(dirName, walkContentDir)
	if err != nil {
		return nil, err
	}

	return &files, nil
}

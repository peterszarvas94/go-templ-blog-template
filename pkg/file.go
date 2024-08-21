package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"peterszarvas94/blog/config"
	"sort"
	"strings"
	"time"
)

type FileData struct {
	Fileroute string      `json:"fileroute"`
	Matter    FrontMatter `json:"matter"`
	DateTime  time.Time   `json:"datetime"`
	Html      string      `json:"html"`
	Content   string      `json:"content"`
	Path      string      `json:"path"`
}

var files []*FileData

func GetFiles() []*FileData {
	return files
}

func CollectFiles() error {
	files = files[:0]

	dirName := config.Dirs.Content

	err := filepath.Walk(dirName, walkContentDir)
	if err != nil {
		return err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].DateTime.After(files[j].DateTime)
	})

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

	fileroute := strings.TrimSuffix(path, ".md")
	fileroute = strings.TrimPrefix(fileroute, config.Dirs.Content)

	dateTime, err := parseDateTime(&matter)
	if err != nil {
		return fmt.Errorf("failed to parse date time: %s", path)
	}

	content := StripHTMLTags(html)

	// Get the first part (before the dot)
	file := &FileData{
		Fileroute: fileroute,
		Matter:    matter,
		DateTime:  dateTime,
		Html:      html,
		Content:   content,
		Path:      path,
	}
	files = append(files, file)

	return nil
}

func FindFileFromFilePath(filePath string) (*FileData, error) {
	// TODO: assert collect files is run (not nil)
	for _, file := range files {
		if file.Path == filePath {
			return file, nil
		}
	}

	return nil, fmt.Errorf("file not found: %s", filePath)
}

func GetTags() map[string][]*FileData {
	tags := make(map[string][]*FileData)

	for _, file := range files {
		for _, tag := range file.Matter.Tags {
			tags[tag] = append(tags[tag], file)
		}
	}

	return tags
}

func GetCategories() map[string][]*FileData {
	categories := make(map[string][]*FileData)
	for _, file := range files {
		categories[file.Matter.Category] = append(categories[file.Matter.Category], file)
	}
	return categories
}

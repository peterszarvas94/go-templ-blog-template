package pkg

import (
	"fmt"
	"strconv"
	"strings"
)

type FileNotFoundError struct {
	message string
}

func (e FileNotFoundError) Error() string {
	return e.message
}

func FindFile(year, month, day, filename string) (*FileData, error) {
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return nil, err
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil {
		return nil, err
	}

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.Date.Year() == yearInt &&
			int(file.Date.Month()) == monthInt &&
			file.Date.Day() == dayInt &&
			strings.TrimSuffix(file.Filename, ".md") == filename {
			return &file, nil
		}
	}

	return nil, FileNotFoundError{
		message: fmt.Sprintf("file not found: %s/%s/%s/%s.md", year, month, day, filename),
	}
}

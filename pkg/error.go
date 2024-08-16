package pkg

import "fmt"

type ProtectedDirError struct {
	dirName string
}

func (e *ProtectedDirError) Error() string {
	return fmt.Sprintf(
		"You can not use \"%s\" as directory name in content folder",
		e.dirName,
	)
}

type FileNotFoundError struct {
	message string
}

func (e FileNotFoundError) Error() string {
	return e.message
}

package pkg

import "fmt"

type ProtectedNameError struct {
	name string
	kind string
}

func (e *ProtectedNameError) Error() string {
	return fmt.Sprintf(
		"You can not use \"%s\" as %sname in content folder",
		e.name,
		e.kind,
	)
}

// type FileNotFoundError struct {
// 	fileName string
// }
//
// func (e *FileNotFoundError) Error() string {
// 	return fmt.Sprintf(
// 		"File \"%s\" not found",
// 		e.fileName,
// 	)
// }

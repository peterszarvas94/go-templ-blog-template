package pkg

import (
	"fmt"
	"path"
)

func Href(file *FileData) string {
	fileName := path.Join("posts", file.Date.Format("2006/01/02"), file.Filename)
	withTrail := fmt.Sprintf("/%s", fileName)
	return withTrail
}

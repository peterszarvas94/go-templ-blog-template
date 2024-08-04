package pkg

import (
	"bytes"
	"path"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
)

func parseFrontMatter(input string) (FrontMatter, string, error) {
	var frontMatter FrontMatter

	body, err := frontmatter.MustParse(strings.NewReader(input), &frontMatter)
	if err != nil {
		return frontMatter, "", err
	}

	return frontMatter, string(body), nil
}

func parseFileContent(input string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(input), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetPath(post *FileInfo) string {
	withMd := path.Join("posts", post.Date.Format("2006/01/02"), post.Filename)
	return withMd[:len(withMd)-3]
}

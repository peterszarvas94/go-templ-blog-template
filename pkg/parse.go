package pkg

import (
	"bytes"
	"path"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-emoji"
	// "github.com/yuin/goldmark-emoji/definition"
	"github.com/yuin/goldmark/extension"
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
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, emoji.Emoji),
	)

	var buf bytes.Buffer
	if err := md.Convert([]byte(input), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetPath(post *FileInfo) string {
	withMd := path.Join("posts", post.Date.Format("2006/01/02"), post.Filename)
	return withMd[:len(withMd)-3]
}

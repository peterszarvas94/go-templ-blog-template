package pkg

import (
	"bytes"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-emoji"

	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
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
	markdown := goldmark.New(
		goldmark.WithExtensions(
			extension.GFM,
			extension.DefinitionList,
			extension.Footnote,
			emoji.Emoji,
		),
		goldmark.WithParserOptions(
			parser.WithAttribute(),
		),
	)

	var buf bytes.Buffer
	if err := markdown.Convert([]byte(input), &buf); err != nil {
		return "", err
	}

	return buf.String(), nil
}

package pages

import (
	"peterszarvas94/blog/pkg/fileutils"
	"peterszarvas94/blog/theme/templates"

	"github.com/a-h/templ"
)

type (
	NotFoundPage func() templ.Component
	IndexPage    func(files []*fileutils.FileData) templ.Component
	PostPage     func(post *fileutils.FileData) templ.Component
	TagPage      func(tag string, files []*fileutils.FileData) templ.Component
	CategoryPage func(category string, files []*fileutils.FileData) templ.Component
)

var (
	NotFound NotFoundPage = templates.NotFoundPage
	Index    IndexPage    = templates.IndexPage
	Post     PostPage     = templates.PostPage
	Tag      TagPage      = templates.TagPage
	Category CategoryPage = templates.CategoryPage
)

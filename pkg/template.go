package pkg

import "github.com/a-h/templ"

type (
	NotFoundPage func() templ.Component
	IndexPage    func(files []*FileData) templ.Component
	PostPage     func(post *FileData) templ.Component
	TagPage      func(tag string, files []*FileData) templ.Component
	CategoryPage func(category string, files []*FileData) templ.Component
)

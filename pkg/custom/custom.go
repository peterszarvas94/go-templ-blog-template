package custom

import (
	"peterszarvas94/blog/pkg/fileutils"
	"peterszarvas94/blog/theme/templates"

	"github.com/a-h/templ"
)

type routeMap map[string]templ.Component

var Routes = &routeMap{
	"/":       templates.CustomIndexPage(fileutils.GetFileByTitle("index")),
	"/search": templates.SearchPage(fileutils.GetFiles()),
	"/docs":   templates.PostsPage(fileutils.GetFiles()),
}

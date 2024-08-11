package pkg

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func HtmlString(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}

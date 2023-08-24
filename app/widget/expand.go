package widget

import (
	"html/template"
)

func Expand(params ...interface{}) template.HTML {
	html := ``
	rt := template.HTML(html)
	return rt
}

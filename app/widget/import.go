package widget

import (
	"html/template"
)

func Import(params ...interface{}) template.HTML {
	html := ``
	rt := template.HTML(html)
	return rt
}

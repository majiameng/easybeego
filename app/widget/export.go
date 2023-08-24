package widget

import (
	"html/template"
)

func Export(params ...interface{}) template.HTML {
	html := ``
	rt := template.HTML(html)
	return rt
}

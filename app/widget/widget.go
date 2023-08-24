package widget

import (
	"html/template"
)

func Widget(params ...interface{}) template.HTML {
	html := ``
	rt := template.HTML(html)
	return rt
}

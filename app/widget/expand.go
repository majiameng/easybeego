package widget

import (
	"html/template"
)

func Expand(params ...interface{}) template.HTML {
	rt := template.HTML(`html`)
	return rt
}

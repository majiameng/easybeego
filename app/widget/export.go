package widget

import (
	"html/template"
)

func Export(params ...interface{}) template.HTML {
	rt := template.HTML(`html`)
	return rt
}

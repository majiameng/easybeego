package widget

import (
	"html/template"
)

func Widget(params ...interface{}) template.HTML {
	rt := template.HTML(`html`)
	return rt
}

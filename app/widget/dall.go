package widget

import (
	"html/template"
)

func Dall(name string) template.HTML {
	html := `<button type="button" class="layui-btn layui-btn-sm" lay-event="dall">` + name + `</button>`
	rt := template.HTML(html)
	return rt
}

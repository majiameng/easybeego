package widget

import (
	"html/template"
)

func Edit(name string) template.HTML {
	html := `<button class="layui-btn layui-btn-sm" lay-event="edit">` + name + `</button>`
	rt := template.HTML(html)
	return rt
}

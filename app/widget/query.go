package widget

import (
	"html/template"
)

func Query(name string) template.HTML {
	html := `<button class="layui-btn layui-btn-sm" lay-event="query">` + name + `</button>`
	rt := template.HTML(html)
	return rt
}

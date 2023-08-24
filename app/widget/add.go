package widget

import (
	"html/template"
)

func Add(name string, params ...interface{}) template.HTML {
	html := `<button type="button" class="layui-btn layui-btn-sm" lay-event="add">` + name + `</button>`
	rt := template.HTML(html)
	return rt
}

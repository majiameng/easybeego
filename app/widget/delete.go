package widget

import (
	"html/template"
)

func Delete(name string) template.HTML {
	html := `<button type="button" class="layui-btn layui-btn-sm" lay-event="del">` + name + `</button>`
	rt := template.HTML(html)
	return rt
}

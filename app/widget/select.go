package widget

import (
	"fmt"
	"html/template"
)

type SelectParams struct {
	from   string
	option string
	value  interface{}
}

func Select(from string, option string, value interface{}) template.HTML {
	fmt.Println("11============")
	fmt.Println(from)
	fmt.Println(option)
	fmt.Println(value)
	rt := template.HTML(`
    <select>
      <option value="">请选择</option>
      <option value="AAA">选项 A</option>
      <option value="BBB">选项 B</option>
      <option value="CCC">选项 C</option>
    </select>
`)
	return rt
}

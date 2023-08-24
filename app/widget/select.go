package widget

import (
	"fmt"
	"html/template"
	"strings"
)

//type SelectParams struct {
//	from   string
//	option string
//	value  interface{}
//}

type SelectParams struct {
	one   string
	two   string
	three string
	four  string
}

/**
from: gender|0|性别|name|id
option: 1=男,2=女,3=保密
value: 0
*/
func Select(from string, option string, value interface{}) template.HTML {
	fmt.Println("11============")
	fromArray := strings.Split(from, "|")
	optionArray := strings.Split(option, ",")
	//表单名称 是否必传 提示问题 name值 id值
	fromName, must, placeholder, name, id := fromArray[0], fromArray[1], fromArray[2], fromArray[3], fromArray[4]
	fmt.Println(name, id)

	html := ""
	html += `
    <select name="` + fromName + `" `
	if must == "1" {
		html += `lay-verify="required"`
	}
	if placeholder != "" {
		html += `placeholder="` + placeholder + `"`
	}
	html += `><option value="">请选择</option>`
	for _, option := range optionArray {
		optionArray := strings.Split(option, "=")
		html += `<option value="` + optionArray[0] + `">` + optionArray[1] + `</option>`
		fmt.Println(optionArray)
	}
	html += `</select>`
	fmt.Println(html)
	rt := template.HTML(html)
	return rt
}

package widget

import (
	"fmt"
	"html/template"
	"strings"
)

type SelectParams struct {
	one   string
	two   string
	three string
	four  string
}

//func witchType(value interface{}) {
//	var t string = ""
//	switch value.(type) {
//	case bool:
//		t = "string"
//	case int:
//		t = "int"
//	case *bool:
//		t = "bool"
//	case *int:
//		t = "int"
//	default:
//		t = "int"
//	}
//	return t
//}

/**
from: gender|0|性别|name|id
option: 1=男,2=女,3=保密
value: 0
*/
func Select(from string, option interface{}, value interface{}) template.HTML {
	fmt.Println("tinymeng:", from, option, value)

	fromArray := strings.Split(from, "|")
	//表单名称 是否必传 提示问题 name值 id值
	fromName, must, placeholder, name, id := fromArray[0], fromArray[1], fromArray[2], fromArray[3], fromArray[4]
	fmt.Println(name, id)

	var optionArray []string
	if v, ok := option.(string); ok {
		fmt.Println(" string ")

		optionArray = strings.Split(v, ",")
		fmt.Println(v)
	} else if v, ok := option.(map[int]string); ok {
		fmt.Println("map int string ")
		fmt.Println(optionArray)
		//optionArray = option
		fmt.Println(v)
	} else {
		fmt.Println("else ")

	}
	fmt.Println(optionArray)

	html := `
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
	html = `</select>`
	fmt.Println(html)
	rt := template.HTML(html)
	return rt
}

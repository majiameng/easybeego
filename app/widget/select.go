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
gender|0|性别|name|id
1=男,2=女,3=保密
0
*/
func Select(from string, option string, value interface{}) template.HTML {
	fmt.Println("11============")
	fmt.Println(from) //gender|0|性别|name|id
	fromData := strings.Split(from, "|")
	fmt.Println(fromData) //[gender 0 性别 name id]

	//var selectParams SelectParams
	//for {
	//
	//}

	//fmt.Println(22222, selectPar)

	//for _, v := range list {
	//	adSortList[v.Id] = v.Description
	//}

	fmt.Println(option)
	fmt.Println(value)

	html := ""
	html += `
    <select>
      <option value="">请选择</option>
      <option value="AAA">选项 A</option>
      <option value="BBB">选项 B</option>
      <option value="CCC">选项 C</option>
    </select>
`
	rt := template.HTML(html)
	return rt
}

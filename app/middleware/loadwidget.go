/**
 * @author: Tinymeng <666@majiameng.com>
 */

package middleware

import (
	"easybeego/app/widget"
	beego "github.com/beego/beego/v2/server/web"
	"html/template"
)

func LoadWidget() {
	// 自定义模板函数
	beego.AddFuncMap("widget", widget.Widget)
	// 查询按钮
	beego.AddFuncMap("query", widget.Query)
	// 添加按钮
	beego.AddFuncMap("add", widget.Add)
	// 编辑按钮
	beego.AddFuncMap("edit", widget.Edit)
	// 删除按钮
	beego.AddFuncMap("delete", widget.Delete)
	// 批量删除按钮
	beego.AddFuncMap("dall", widget.Dall)
	// 展开按钮
	beego.AddFuncMap("expand", widget.Expand)
	// 收缩按钮
	beego.AddFuncMap("collapse", widget.Collapse)
	// 添加子级按钮
	beego.AddFuncMap("addz", widget.Addz)
	// 开关按钮
	beego.AddFuncMap("switch", widget.Switch)
	// 提交按钮
	beego.AddFuncMap("submit", widget.Submit)
	// 导入按钮
	beego.AddFuncMap("import", widget.Import)
	// 导出按钮
	beego.AddFuncMap("export", widget.Export)
	// 图标选择
	beego.AddFuncMap("icon", widget.Icon)
	// 选择下拉组件
	beego.AddFuncMap("select", widget.Select)
	// 穿梭组件
	beego.AddFuncMap("transfer", widget.Transfer)
	// 上传单图
	beego.AddFuncMap("upload_image", widget.UploadImage)
	// 上传图集
	beego.AddFuncMap("album", widget.Album)
	// 复选框组件
	beego.AddFuncMap("checkbox", widget.Checkbox)
	// 单选按钮组价
	beego.AddFuncMap("radio", widget.Radio)
	// 城市组件
	beego.AddFuncMap("city", widget.City)
	// 日期选择组件
	beego.AddFuncMap("date", widget.Date)
	// 富文本组件
	beego.AddFuncMap("kindeditor", widget.Kindeditor)
	// 站点
	beego.AddFuncMap("item", widget.Item)
	// 自定义组件
	beego.AddFuncMap("safe", func(str string) template.HTML {
		return template.HTML(str)
	})
}

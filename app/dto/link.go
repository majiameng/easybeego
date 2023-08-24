/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询条件
type LinkPageReq struct {
	Name     string `form:"name"`     // 友链名称
	Type     int    `form:"type"`     // 友链类型
	Platform int    `form:"platform"` // 投放平台
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加友链
type LinkAddReq struct {
	Name     string `form:"name" validate:"required"` // 友链名称
	Type     int    `form:"type" validate:"int"`      // 类型：1友情链接 2合作伙伴
	Url      string `form:"url"`                      // 友链地址
	ItemId   int    `form:"itemId"`                   // 站点ID
	CateId   int    `form:"cateId"`                   // 栏目ID
	Platform int    `form:"platform" validate:"int"`  // 平台：1PC站 2WAP站 3微信小程序 4APP应用
	Form     int    `form:"form" validate:"int"`      // 友链形式：1文字链接 2图片链接
	Image    string `form:"image"`                    // 友链图片
	Status   int    `form:"status" validate:"int"`    // 状态：1在用 2停用
	Sort     int    `form:"sort" validate:"int"`      // 显示顺序
	Note     string `form:"note"`                     // 备注
}

// 添加友链表单验证
func (v LinkAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "友链名称不能为空.",
		"Type.int":      "请选择友链类型.",
		"Platform.int":  "请选择友链平台.",
		"Form.int":      "请选择友链形式.",
		"Status.int":    "请选择友链状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 修改友链
type LinkUpdateReq struct {
	Id       int    `form:"id" validate:"int"`
	Name     string `form:"name" validate:"required"` // 友链名称
	Type     int    `form:"type" validate:"int"`      // 类型：1友情链接 2合作伙伴
	Url      string `form:"url"`                      // 友链地址
	ItemId   int    `form:"itemId"`                   // 站点ID
	CateId   int    `form:"cateId"`                   // 栏目ID
	Platform int    `form:"platform" validate:"int"`  // 平台：1PC站 2WAP站 3微信小程序 4APP应用
	Form     int    `form:"form" validate:"int"`      // 友链形式：1文字链接 2图片链接
	Image    string `form:"image"`                    // 友链图片
	Status   int    `form:"status" validate:"int"`    // 状态：1在用 2停用
	Sort     int    `form:"sort" validate:"int"`      // 显示顺序
	Note     string `form:"note"`                     // 备注
}

// 更新友链表单验证
func (v LinkUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "友链ID不能为空.",
		"Name.required": "友链名称不能为空.",
		"Type.int":      "请选择友链类型.",
		"Platform.int":  "请选择友链平台.",
		"Form.int":      "请选择友链形式.",
		"Status.int":    "请选择友链状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 设置状态
type LinkStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

func (v LinkStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "友链ID不能为空.",
		"Status.int": "请选择友链状态.",
	}
}

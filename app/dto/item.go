/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询条件
type ItemPageReq struct {
	Name  string `form:"name"`  // 站点名称
	Type  int    `form:"type"`  // 站点类型
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加站点
type ItemAddReq struct {
	Name   string `form:"name" validate:"required"` // 站点名称
	Type   int    `form:"type" validate:"int"`      // 站点类型:1普通站点 2其他
	Url    string `form:"url" validate:"required"`  // 站点地址
	Image  string `form:"image"`                    // 站点图片
	Status int    `form:"status" validate:"int"`    // 状态：1在用 2停用
	Note   string `form:"note"`                     // 站点备注
	Sort   int    `form:"sort" validate:"int"`      // 显示顺序
}

// 添加站点表单验证
func (v ItemAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "站点名称不能为空.",
		"Type.int":      "请选择站点类型.",
		"Url.required":  "站点地址不能为空.",
		"Status.int":    "请选择站点状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新站点
type ItemUpdateReq struct {
	Id     int    `form:"id" validate:"int"`
	Name   string `form:"name" validate:"required"` // 站点名称
	Type   int    `form:"type" validate:"int"`      // 站点类型:1普通站点 2其他
	Url    string `form:"url" validate:"required"`  // 站点地址
	Image  string `form:"image"`                    // 站点图片
	Status int    `form:"status" validate:"int"`    // 状态：1在用 2停用
	Note   string `form:"note"`                     // 站点备注
	Sort   int    `form:"sort" validate:"int"`      // 显示顺序
}

// 更新站点表单验证
func (v ItemUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "站点ID不能为空.",
		"Name.required": "站点名称不能为空.",
		"Type.int":      "请选择站点类型.",
		"Url.required":  "站点地址不能为空.",
		"Status.int":    "请选择站点状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 设置状态
type ItemStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

// 设置状态参数验证
func (v ItemStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "站点ID不能为空.",
		"Status.int": "请选择站点状态.",
	}
}

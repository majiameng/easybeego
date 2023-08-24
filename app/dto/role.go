/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询条件
type RolePageReq struct {
	Name  string `form:"name"`  // 角色名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加角色
type RoleAddReq struct {
	Name   string `form:"name" validate:"required"`
	Code   string `form:"code" validate:"required"`
	Status int    `form:"status" validate:"int"`
	Sort   int    `form:"sort" validate:"int"`
	Note   string `form:"note"`
}

// 添加角色表单验证
func (v RoleAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "角色名称不能为空.",
		"Code.required": "角色编码不能为空.",
		"Status.int":    "请选择角色状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新角色
type RoleUpdateReq struct {
	Id     int    `form:"id" validate:"int"`
	Name   string `form:"name" validate:"required"`
	Code   string `form:"code" validate:"required"`
	Status int    `form:"status" validate:"int"`
	Sort   int    `form:"sort" validate:"int"`
	Note   string `form:"note"`
}

// 添加角色表单验证
func (v RoleUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "角色ID不能为空.",
		"Name.required": "角色名称不能为空.",
		"Code.required": "角色编码不能为空.",
		"Status.int":    "请选择角色状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 设置状态
type RoleStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

// 设置状态验证
func (v RoleStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "角色ID不能为空.",
		"Status.int": "请选择角色状态.",
	}
}

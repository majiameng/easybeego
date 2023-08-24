/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询条件
type DeptPageReq struct {
	Name  string `form:"name"`  // 部门名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加部门
type DeptAddReq struct {
	Name     string `form:"name" validate:"required"`
	Code     string `form:"code" validate:"required"`
	Fullname string `form:"fullname" validate:"required"`
	Type     int    `form:"type" validate:"int"`
	Pid      int    `form:"pid" validate:"int"`
	Sort     int    `form:"sort" validate:"int"`
	Note     string
}

// 添加部门表单验证
func (v DeptAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required":     "部门名称不能为空.",
		"Code.required":     "部门编码不能为空.",
		"Fullname.required": "部门全称不能为空.",
		"Type.int":          "请选择部门类型.",
		"Sort.int":          "排序不能为空.",
	}
}

// 更新部门
type DeptUpdateReq struct {
	Id       int    `form:"id" validate:"int"`
	Name     string `form:"name" validate:"required"`
	Code     string `form:"code" validate:"required"`
	Fullname string `form:"fullname" validate:"required"`
	Type     int    `form:"type" validate:"int"`
	Pid      int    `form:"pid" validate:"int"`
	Sort     int    `form:"sort" validate:"int"`
	Note     string
}

// 添加部门表单验证
func (v DeptUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":            "部门ID不能为空.",
		"Name.required":     "部门名称不能为空.",
		"Code.required":     "部门编码不能为空.",
		"Fullname.required": "部门全称不能为空.",
		"Type.int":          "请选择部门类型.",
		"Sort.int":          "排序不能为空.",
	}
}

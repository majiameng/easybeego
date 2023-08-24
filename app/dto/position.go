/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询条件
type PositionPageReq struct {
	Name  string `form:"name"`  // 岗位名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加岗位
type PositionAddReq struct {
	Name   string `form:"name" validate:"required"`
	Status int    `form:"status" validate:"int"`
	Sort   int    `form:"sort" validate:"int"`
}

// 添加岗位表单验证
func (v PositionAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "岗位名称不能为空.",
		"Status.int":    "请选择岗位状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新岗位
type PositionUpdateReq struct {
	Id     int    `form:"id" validate:"int"`
	Name   string `form:"name" validate:"required"`
	Status int    `form:"status" validate:"int"`
	Sort   int    `form:"sort" validate:"int"`
}

// 更新岗位表单验证
func (v PositionUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "岗位ID不能为空.",
		"Name.required": "岗位名称不能为空.",
		"Status.int":    "请选择岗位状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 设置状态
type PositionStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status"    validate:"int"`
}

func (v PositionStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "岗位ID不能为空.",
		"Status.int": "请选择岗位状态.",
	}
}

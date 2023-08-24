/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询
type LevelPageReq struct {
	Name  string `form:"name"`  // 职级名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加职级
type LevelAddReq struct {
	Name   string `form:"name"  validate:"required"`
	Status int    `form:"status"    validate:"int"`
	Sort   int    `form:"sort" validate:"int"`
}

// 添加职级表单验证
func (v LevelAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "职级名称不能为空.",
		"Status.int":    "请选择职级状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新职级
type LevelUpdateReq struct {
	Id     int    `form:"id" validate:"int"`
	Name   string `form:"name"  validate:"required"`
	Status int    `form:"status"    validate:"int"`
	Sort   int    `form:"sort" validate:"int"`
}

// 更新职级表单验证
func (v LevelUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "职级ID不能为空.",
		"Name.required": "职级名称不能为空.",
		"Status.int":    "请选择职级状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 设置状态
type LevelStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status"    validate:"int"`
}

// 设置状态参数验证
func (v LevelStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "职级ID不能为空.",
		"Status.int": "请选择职级状态.",
	}
}

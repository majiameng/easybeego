/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询
type ConfigPageReq struct {
	Name  string `form:"name"`  // 配置名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加配置
type ConfigAddReq struct {
	Name string `form:"Name" validate:"required"` // 配置名称
	Sort int    `form:"Sort" validate:"int"`      // 显示顺序
}

// 添加配置表单验证
func (v ConfigAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "配置名称不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

// 修改配置
type ConfigUpdateReq struct {
	Id   int    `form:"Id" validate:"int"`        // 主键ID
	Name string `form:"Name" validate:"required"` // 配置名称
	Sort int    `form:"Sort" validate:"int"`      // 显示顺序
}

// 修改配置表单验证
func (v ConfigUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "配置ID不能为空.",
		"Name.required": "配置名称不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

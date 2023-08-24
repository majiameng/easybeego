/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询
type DictPageReq struct {
	Name  string `form:"name"`  // 字典名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加字典
type DictAddReq struct {
	Name string `form:"Name" validate:"required"` // 字典名称
	Code string `form:"Code" validate:"required"` // 字典值
	Sort int    `form:"Sort" validate:"int"`      // 显示顺序
	Note string `form:"Note"`                     // 字典备注
}

// 添加字典表单验证
func (v DictAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "字典名称不能为空.",
		"Code.required": "字典编码不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

// 修改字典
type DictUpdateReq struct {
	Id   int    `form:"Id" validate:"int"`        // 主键ID
	Name string `form:"Name" validate:"required"` // 字典名称
	Code string `form:"Code" validate:"required"` // 字典值
	Sort int    `form:"Sort" validate:"int"`      // 显示顺序
	Note string `form:"Note"`                     // 字典备注
}

// 更新字典表单验证
func (v DictUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "字典ID不能为空.",
		"Name.required": "字典名称不能为空.",
		"Code.required": "字典编码不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

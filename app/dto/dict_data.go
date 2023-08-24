/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 字典项列表查询条件
type DictDataPageReq struct {
	DictId int    `form:"dictId"` // 字典ID
	Name   string `form:"name"`   // 字典项名称
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
}

// 添加字典项
type DictDataAddReq struct {
	Name   string `form:"Name,unique" validate:"required"` // 字典项名称
	Code   string `form:"Code" validate:"required"`        // 字典项值
	DictId int    `form:"DictId" validate:"int"`           // 字典类型ID
	Note   string `form:"Note"`                            // 备注
	Sort   int    `form:"Sort" validate:"int"`             // 显示顺序
}

// 添加字典项表单验证
func (v DictDataAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "字典项名称不能为空.",
		"Code.required": "字典项编码不能为空.",
		"DictId.int":    "请选择数据字典.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新字典项
type DictDataUpdateReq struct {
	Id     int    `form:"Id" validate:"int"`
	Name   string `form:"Name,unique" validate:"required"` // 字典项名称
	Code   string `form:"Code" validate:"required"`        // 字典项值
	DictId int    `form:"DictId" validate:"int"`           // 字典类型ID
	Note   string `form:"Note"`                            // 备注
	Sort   int    `form:"Sort" validate:"int"`             // 显示顺序
}

// 更新字典项表单验证
func (v DictDataUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "字典项ID不能为空.",
		"Name.required": "字典项名称不能为空.",
		"Code.required": "字典项编码不能为空.",
		"DictId.int":    "请选择数据字典.",
		"Sort.int":      "排序不能为空.",
	}
}

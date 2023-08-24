/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 字典项列表查询条件
type ConfigDataPageReq struct {
	ConfigId int    `form:"configId"` // 字典ID
	Title    string `form:"name"`     // 配置标题
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加字典项
type ConfigDataAddReq struct {
	Title    string `form:"Title" validate:"required"` // 配置标题
	Code     string `form:"Code" validate:"required"`  // 配置编码
	Value    string `form:"Value"`                     // 配置值
	Options  string `form:"Options"`                   // 配置项
	ConfigId int    `form:"ConfigId" validate:"int"`   // 配置ID
	Type     string `form:"Type" validate:"required"`  // 配置类型
	Sort     int    `form:"Sort" validate:"int"`       // 排序
	Note     string `form:"Note"`                      // 配置说明
}

// 添加配置项表单验证
func (v ConfigDataAddReq) Messages() map[string]string {
	return validate.MS{
		"Title.required": "配置项名称不能为空.",
		"Code.required":  "配置项编码不能为空.",
		"ConfigId.int":   "配置ID不能为空.",
		"Sort.int":       "排序不能为空.",
	}
}

// 更新字典项
type ConfigDataUpdateReq struct {
	Id       int    `form:"Id" validate:"int"`
	Title    string `form:"Title" validate:"required"` // 配置标题
	Code     string `form:"Code" validate:"required"`  // 配置编码
	Value    string `form:"Value"`                     // 配置值
	Options  string `form:"Options"`                   // 配置项
	ConfigId int    `form:"ConfigId" validate:"int"`   // 配置ID
	Type     string `form:"Type" validate:"required"`  // 配置类型
	Sort     int    `form:"Sort" validate:"int"`       // 排序
	Note     string `form:"Note"`                      // 配置说明
}

// 更新配置项表单验证
func (v ConfigDataUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":         "配置项ID不能为空.",
		"Title.required": "配置项名称不能为空.",
		"Code.required":  "配置项编码不能为空.",
		"ConfigId.int":   "配置ID不能为空.",
		"Sort.int":       "排序不能为空.",
	}
}

// 设置状态
type ConfigDataStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

// 设置状态参数验证
func (v ConfigDataStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "配置项ID不能为空.",
		"Status.int": "请选择配置项状态.",
	}
}

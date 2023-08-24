/**
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 演示一Dto
 * @author 半城风雨
 * @since 2022-05-13
 * @File : example
 */
package dto

import "github.com/gookit/validate"

// 分页查询
type ExamplePageReq struct {
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
	Name   string `form:"name"`   // 测试名称
	Status int    `form:"status"` // 状态：1正常 2停用
	Type   int    `form:"type"`   // 类型：1京东 2淘宝 3拼多多 4唯品会
	IsVip  int    `form:"isVip"`  // 是否VIP：1是 2否
}

// 添加演示一
type ExampleAddReq struct {
	Name    string `form:"name" validate:"required"`    // 测试名称
	Avatar  string `form:"avatar" validate:"required"`  // 头像
	Content string `form:"content" validate:"required"` // 内容
	Status  int    `form:"status" validate:"int"`       // 状态：1正常 2停用
	Type    int    `form:"type" validate:"int"`         // 类型：1京东 2淘宝 3拼多多 4唯品会
	IsVip   int    `form:"isVip" validate:"int"`        // 是否VIP：1是 2否
	Sort    int    `form:"sort" validate:"int"`         // 排序号
}

// 添加表单验证
func (v ExampleAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required":    "测试名称不能为空.", // 测试名称
		"Avatar.required":  "头像不能为空.",   // 头像
		"Content.required": "内容不能为空.",   // 内容
		"Status.int":       "请选择状态.",    // 状态：1正常 2停用
		"Type.int":         "请选择类型.",    // 类型：1京东 2淘宝 3拼多多 4唯品会
		"IsVip.int":        "请选择是否VIP.", // 是否VIP：1是 2否
		"Sort.int":         "排序号不能为空.",  // 排序号
	}
}

// 编辑演示一
type ExampleUpdateReq struct {
	Id      int    `form:"id" validate:"int"`
	Name    string `form:"name" validate:"required"`    // 测试名称
	Avatar  string `form:"avatar" validate:"required"`  // 头像
	Content string `form:"content" validate:"required"` // 内容
	Status  int    `form:"status" validate:"int"`       // 状态：1正常 2停用
	Type    int    `form:"type" validate:"int"`         // 类型：1京东 2淘宝 3拼多多 4唯品会
	IsVip   int    `form:"isVip" validate:"int"`        // 是否VIP：1是 2否
	Sort    int    `form:"sort" validate:"int"`         // 排序号
}

// 更新表单验证
func (v ExampleUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":           "记录ID不能为空.",
		"Name.required":    "测试名称不能为空.", // 测试名称
		"Avatar.required":  "头像不能为空.",   // 头像
		"Content.required": "内容不能为空.",   // 内容
		"Status.int":       "请选择状态.",    // 状态：1正常 2停用
		"Type.int":         "请选择类型.",    // 类型：1京东 2淘宝 3拼多多 4唯品会
		"IsVip.int":        "请选择是否VIP.", // 是否VIP：1是 2否
		"Sort.int":         "排序号不能为空.",  // 排序号
	}
}

// 设置状态
type ExampleStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

// 设置状态参数验证
func (v ExampleStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "记录ID不能为空.",
		"Status.int": "请选择状态：1正常 2停用.",
	}
}

// 设置是否VIP
type ExampleIsVipReq struct {
	Id    int `form:"id" validate:"int"`
	IsVip int `form:"isVip" validate:"int"`
}

// 设置状态参数验证
func (v ExampleIsVipReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":    "记录ID不能为空.",
		"IsVip.int": "请选择是否VIP：1是 2否.",
	}
}

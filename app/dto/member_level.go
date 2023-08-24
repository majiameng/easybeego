/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 查询会员等级
type MemberLevelPageReq struct {
	Name  string `form:"name"`  // 等级名称
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加会员等级
type MemberLevelAddReq struct {
	Name string `form:"name" validate:"required"` // 级别名称
	Sort int    `form:"sort" validate:"int"`      // 排序号
}

// 添加会员等级表单验证
func (v MemberLevelAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "会员等级名称不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

// 更新会员等级
type MemberLevelUpdateReq struct {
	Id   int    `form:"id" validate:"int"`
	Name string `form:"name" validate:"required"` // 级别名称
	Sort int    `form:"sort" validate:"int"`      // 排序号
}

// 添加会员等级表单验证
func (v MemberLevelUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "会员等级ID不能为空.",
		"Name.required": "会员等级名称不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

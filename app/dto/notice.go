/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询
type NoticePageReq struct {
	Title  string `form:"title"`  // 通知标题
	Source int    `form:"source"` // 通知来源
	Page   int    `form:"page"`   // 页码
	Limit  int    `form:"limit"`  // 每页数
}

// 添加通知公告
type NoticeAddReq struct {
	Title   string `form:"title" validate:"required"`   // 通知标题
	Content string `form:"content" validate:"required"` // 通知内容
	Source  int    `form:"source" validate:"int"`       // 来源：1内部通知 2外部新闻
	IsTop   int    `form:"isTop" validate:"int"`        // 是否置顶：1是 2否
	Status  int    `form:"status" validate:"int"`       // 状态：1已发布 2待发布
	Sort    int    `form:"sort" validate:"int"`         // 排序号
}

// 添加通知表单验证
func (v NoticeAddReq) Messages() map[string]string {
	return validate.MS{
		"Title.required":   "通知标题不能为空.",
		"Content.required": "通知内容不能为空.",
		"Source.int":       "请选择通知来源.",
		"IsTop.int":        "请选择是否置顶.",
		"Status.int":       "请选择通知状态.",
		"Sort.int":         "排序号不能为空.",
	}
}

// 更新通知公告
type NoticeUpdateReq struct {
	Id      int    `form:"id" validate:"int"`
	Title   string `form:"title" validate:"required"`   // 通知标题
	Content string `form:"content" validate:"required"` // 通知内容
	Source  int    `form:"source" validate:"int"`       // 来源：1内部通知 2外部新闻
	IsTop   int    `form:"isTop" validate:"int"`        // 是否置顶：1是 2否
	Status  int    `form:"status" validate:"int"`       // 状态：1已发布 2待发布
	Sort    int    `form:"sort" validate:"int"`         // 排序号
}

// 添加通知表单验证
func (v NoticeUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":           "通知ID不能为空.",
		"Title.required":   "通知标题不能为空.",
		"Content.required": "通知内容不能为空.",
		"Source.int":       "请选择通知来源.",
		"IsTop.int":        "请选择是否置顶.",
		"Status.int":       "请选择通知状态.",
		"Sort.int":         "排序号不能为空.",
	}
}

// 设置状态
type NoticeStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

// 设置状态参数验证
func (v NoticeStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "通知ID不能为空.",
		"Status.int": "请选择通知状态.",
	}
}

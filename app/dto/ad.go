/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 列表查询
type AdPageReq struct {
	Title string `form:"title"` // 广告标题
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数
}

// 添加广告
type AdAddReq struct {
	Title       string `form:"title" validate:"required"`       // 广告标题
	AdSortId    int    `form:"adSortId" validate:"int"`         // 广告位ID
	Cover       string `form:"cover"`                           // 广告图片
	Type        int    `form:"type" validate:"int"`             // 广告格式：1图片 2文字 3视频 4推荐
	Description string `form:"description" validate:"required"` // 广告描述
	Content     string `form:"content"`                         // 广告内容
	Url         string `form:"url" validate:"required"`         // 广告链接
	Width       int    `form:"width" validate:"int"`            // 广告宽度
	Height      int    `form:"height" validate:"int"`           // 广告高度
	StartTime   string `form:"startTime" validate:"required"`   // 开始时间
	EndTime     string `form:"endTime" validate:"required"`     // 结束时间
	Status      int    `form:"status" validate:"int"`           // 状态：1在用 2停用
	Sort        int    `form:"sort" validate:"int"`             // 排序
	Note        string `form:"note"`                            // 备注
}

// 添加广告表单验证
func (v AdAddReq) Messages() map[string]string {
	return validate.MS{
		"Title.required":       "广告标题不能为空.",
		"AdSortId.int":         "请选择广告位.",
		"Type.int":             "请选择广告格式.",
		"Description.required": "广告描述不能为空.",
		"Url.required":         "广告URL不能为空.",
		"Width.int":            "广告宽度不能为空.",
		"Height.int":           "广告高度不能为空.",
		"StartTime.required":   "开始时间不能为空.",
		"EndTime.required":     "结束时间不能为空.",
		"Status.int":           "请选择广告状态.",
		"Sort.int":             "排序不能为空.",
	}
}

// 更新广告
type AdUpdateReq struct {
	Id          int    `form:"id" validate:"int"`
	Title       string `form:"title" validate:"required"`       // 广告标题
	AdSortId    int    `form:"adSortId" validate:"int"`         // 广告位ID
	Cover       string `form:"cover"`                           // 广告图片
	Type        int    `form:"type" validate:"int"`             // 广告格式：1图片 2文字 3视频 4推荐
	Description string `form:"description" validate:"required"` // 广告描述
	Content     string `form:"content"`                         // 广告内容
	Url         string `form:"url" validate:"required"`         // 广告链接
	Width       int    `form:"width" validate:"int"`            // 广告宽度
	Height      int    `form:"height" validate:"int"`           // 广告高度
	StartTime   string `form:"startTime" validate:"required"`   // 开始时间
	EndTime     string `form:"endTime" validate:"required"`     // 结束时间
	Status      int    `form:"status" validate:"int"`           // 状态：1在用 2停用
	Sort        int    `form:"sort" validate:"int"`             // 排序
	Note        string `form:"note"`                            // 备注
}

// 添加广告表单验证
func (v AdUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":               "广告ID不能为空.",
		"Title.required":       "广告标题不能为空.",
		"AdSortId.int":         "请选择广告位.",
		"Type.int":             "请选择广告格式.",
		"Description.required": "广告描述不能为空.",
		"Url.required":         "广告URL不能为空.",
		"Width.int":            "广告宽度不能为空.",
		"Height.int":           "广告高度不能为空.",
		"StartTime.required":   "开始时间不能为空.",
		"EndTime.required":     "结束时间不能为空.",
		"Status.int":           "请选择广告状态.",
		"Sort.int":             "排序不能为空.",
	}
}

// 设置状态
type AdStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

func (v AdStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "广告ID不能为空.",
		"Status.int": "请选择广告状态.",
	}
}

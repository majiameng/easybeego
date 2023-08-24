/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 列表查询条件
type AdSortPageReq struct {
	Description string `form:"description"` // 广告位描述
	Page        int    `form:"page"`        // 页码
	Limit       int    `form:"limit"`       // 每页数
}

// 添加广告位
type AdSortAddReq struct {
	Description string `form:"description" validate:"required"` // 广告位描述
	ItemId      int    `form:"itemId" validate:"int"`           // 站点ID
	CateId      int    `form:"cateId" validate:"int"`           // 广告位ID
	LocId       int    `form:"locId" validate:"int"`            // 广告页面位置
	Platform    int    `form:"platform" validate:"int"`         // 站点类型：1PC网站 2WAP手机站 3微信小程序 4APP移动端
	Sort        int    `form:"sort"`                            // 广告位排序
}

// 添加广告位表单验证
func (v AdSortAddReq) Messages() map[string]string {
	return validate.MS{
		"Description.required": "广告位描述不能为空.",
		"ItemId.int":           "请选择站点.",
		"CateId.int":           "请选择广告位.",
		"LocId.int":            "位置不能为空.",
		"Platform.int":         "请选择投放平台.",
		"Sort.int":             "排序不能为空.",
	}
}

// 更新广告位
type AdSortUpdateReq struct {
	Id          int    `form:"id" validate:"int"`
	Description string `form:"description" validate:"required"` // 广告位描述
	ItemId      int    `form:"itemId" validate:"int"`           // 站点ID
	CateId      int    `form:"cateId" validate:"int"`           // 广告位ID
	LocId       int    `form:"locId" validate:"int"`            // 广告页面位置
	Platform    int    `form:"platform" validate:"int"`         // 站点类型：1PC网站 2WAP手机站 3微信小程序 4APP移动端
	Sort        int    `form:"sort"`                            // 广告位排序
}

// 添加广告位表单验证
func (v AdSortUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":               "广告位ID不能为空.",
		"Description.required": "广告位描述不能为空.",
		"ItemId.int":           "请选择站点.",
		"CateId.int":           "请选择广告位.",
		"LocId.int":            "位置不能为空.",
		"Platform.int":         "请选择投放平台.",
		"Sort.int":             "排序不能为空.",
	}
}

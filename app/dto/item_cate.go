/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 栏目查询条件
type ItemCateQueryReq struct {
	Name string `form:"name"` // 栏目名称
}

// 添加栏目
type ItemCateAddReq struct {
	Name    string `form:"name" validate:"required"`   // 栏目名称
	Pid     int    `form:"pid" validate:"int"`         // 父级ID
	ItemId  int    `form:"itemId" validate:"int"`      // 栏目ID
	Pinyin  string `form:"pinyin" validate:"required"` // 拼音(全)
	Code    string `form:"code" validate:"required"`   // 拼音(简)
	IsCover int    `form:"isCover" validate:"int"`     // 是否有封面：1是 2否
	Cover   string `form:"cover"`                      // 封面
	Status  int    `form:"status" validate:"int"`      // 状态：1启用 2停用
	Note    string `form:"note"`                       // 备注
	Sort    int    `form:"sort" validate:"int"`        // 排序
}

// 添加栏目表单验证
func (v ItemCateAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required":   "栏目名称不能为空.",
		"Pid.int":         "请选择上级栏目.",
		"ItemId.int":      "请选择栏目ID.",
		"Pinyin.required": "拼音全拼不能为空.",
		"Code.required":   "拼音简拼不能为空.",
		"IsCover.int":     "请选择是否有封面.",
		"Status.int":      "请选择栏目状态.",
		"Sort.int":        "排序不能为空.",
	}
}

// 修改栏目
type ItemCateUpdateReq struct {
	Id      int    `form:"id" validate:"int"`
	Name    string `form:"name" validate:"required"`   // 栏目名称
	Pid     int    `form:"pid" validate:"int"`         // 父级ID
	ItemId  int    `form:"itemId" validate:"int"`      // 栏目ID
	Pinyin  string `form:"pinyin" validate:"required"` // 拼音(全)
	Code    string `form:"code" validate:"required"`   // 拼音(简)
	IsCover int    `form:"isCover" validate:"int"`     // 是否有封面：1是 2否
	Cover   string `form:"cover"`                      // 封面
	Status  int    `form:"status" validate:"int"`      // 状态：1启用 2停用
	Note    string `form:"note"`                       // 备注
	Sort    int    `form:"sort" validate:"int"`        // 排序
}

// 更新栏目表单验证
func (v ItemCateUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":          "栏目ID不能为空.",
		"Name.required":   "栏目名称不能为空.",
		"Pid.int":         "请选择上级栏目.",
		"ItemId.int":      "请选择栏目ID.",
		"Pinyin.required": "拼音全拼不能为空.",
		"Code.required":   "拼音简拼不能为空.",
		"IsCover.int":     "请选择是否有封面.",
		"Status.int":      "请选择栏目状态.",
		"Sort.int":        "排序不能为空.",
	}
}

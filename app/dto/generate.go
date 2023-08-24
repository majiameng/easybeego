/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询条件
type GeneratePageReq struct {
	Name    string `form:"name"`    // 表名称
	Comment string `form:"comment"` // 表描述
	Page    int    `form:"page"`    // 页码
	Limit   int    `form:"limit"`   // 每页数
}

// 单个生成文件
type GenerateFileReq struct {
	Name    string `form:"name" validate:"required"`    // 表名称
	Comment string `form:"comment" validate:"required"` // 表描述
}

// 生成器参数验证
func (r GenerateFileReq) Messages() map[string]string {
	return validate.MS{
		"Name.required":    "数据表名称不能为空.",
		"Comment.required": "数据表描述不能为空.",
	}
}

// 批量生成文件
type BatchGenerateFileReq struct {
	Tables string `form:"tables" validate:"required"` // 表名称
}

// 批量生成器参数验证
func (r BatchGenerateFileReq) Messages() map[string]string {
	return validate.MS{
		"Tables.required": "数据表不能为空.",
	}
}

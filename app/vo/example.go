/**
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 演示一Vo
 * @author 半城风雨
 * @since 2022-05-13
 * @File : example
 */
package vo

import "easybeego/app/models"

// 演示一信息Vo
type ExampleInfoVo struct {
	models.Example

	StatusName int `json:"statusName"` // 状态名称
	TypeName   int `json:"typeName"`   // 类型名称
	IsVipName  int `json:"isVipName"`  // 是否VIP名称
}

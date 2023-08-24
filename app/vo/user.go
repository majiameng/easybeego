/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 用户信息Vo
type UserInfoVo struct {
	models.User
	Birthday     string      `json:"Birthday"`     // 性别
	GenderName   string      `json:"GenderName"`   // 性别
	LevelName    string      `json:"LevelName"`    // 职级
	PositionName string      `json:"PositionName"` // 岗位
	DeptName     string      `json:"DeptName"`     // 部门
	RoleIds      interface{} `json:"RoleIds"`      // 角色ID
	RoleList     interface{} `json:"RoleList"`     // 角色列表
	City         interface{} `json:"City"`         // 省市区
}

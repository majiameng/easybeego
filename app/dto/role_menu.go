/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 角色菜单数据
type RoleMenuSaveReq struct {
	RoleId  int    `form:"roleId" validate:"int"`
	MenuIds string `form:"menuIds"`
}

func (v RoleMenuSaveReq) Messages() map[string]string {
	return validate.MS{
		"RoleId.int": "角色ID不能为空.",
	}
}

/**
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/dto"
	"easybeego/app/services"
	"easybeego/utils/common"
	"github.com/gookit/validate"
)

var RoleMenu = new(RoleMenuController)

type RoleMenuController struct {
	BaseController
}

func (ctl *RoleMenuController) Index() {
	// 角色ID
	roleId, _ := ctl.GetInt("roleId", 0)
	if roleId <= 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "角色ID不能为空",
		})
	}

	// 获取角色菜单权限列表
	list, err := services.RoleMenu.GetRoleMenuList(roleId)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Data: list,
		Msg:  "查询成功",
	})
}

func (ctl *RoleMenuController) Save() {
	// 参数对象
	var req dto.RoleMenuSaveReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 参数验证
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}
	// 调用保存角色菜单数据
	err := services.RoleMenu.Save(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 保存成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "保存成功",
	})
}

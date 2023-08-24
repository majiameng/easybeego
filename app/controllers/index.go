/**
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/dto"
	"easybeego/app/services"
	"easybeego/utils"
	"easybeego/utils/common"
	"github.com/gookit/validate"
	"net/http"
)

var Index = new(IndexController)

type IndexController struct {
	BaseController
}

func (ctl *IndexController) Index() {
	// 获取用户信息
	userInfo := services.Login.GetProfile(utils.Uid(ctl.Ctx))
	// 获取菜单列表
	menuList := services.Menu.GetPermissionMenuList(userInfo.Id)
	ctl.Data["userInfo"] = userInfo
	ctl.Data["menuList"] = menuList
	// 渲染模板
	ctl.TplName = "index.html"
}

func (ctl *IndexController) Main() {
	// 渲染模板
	ctl.TplName = "welcome.html"
}

func (ctl *IndexController) UserInfo() {
	if ctl.Ctx.Input.IsPost() {
		// 参数验证
		var req dto.UserInfoReq
		// 参数绑定
		if err := ctl.ParseForm(&req); err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 参数校验
		v := validate.Struct(req)
		if !v.Validate() {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  v.Errors.One(),
			})
		}
		// 更新信息
		_, err := services.User.UpdateUserInfo(req, utils.Uid(ctl.Ctx))
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 返回结果
		ctl.JSON(common.JsonResult{
			Code: 0,
			Msg:  "更新成功",
		})
	}
	// 获取用户信息
	userInfo := services.Login.GetProfile(utils.Uid(ctl.Ctx))
	// 渲染模板
	ctl.Data["userInfo"] = userInfo
	ctl.TplName = "user/user_info.html"
}

func (ctl *IndexController) UpdatePwd() {
	// 更新密码对象
	var req dto.UpdatePwd
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 参数校验
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}

	// 调用更新密码方法
	rows, err := services.User.UpdatePwd(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "更新密码成功",
	})
}

// 注销系统
func (ctl *IndexController) Logout() {
	// 删除登录Session
	ctl.DelSession("userId")
	// 跳转登录页,方式：301(永久移动),308(永久重定向),307(临时重定向)
	ctl.Redirect("/login", http.StatusFound)
}

/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/constant"
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/app/services"
	"easybeego/utils"
	"easybeego/utils/common"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gookit/validate"
)

var Menu = new(MenuController)

type MenuController struct {
	BaseController
}

func (ctl *MenuController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "menu/index.html"
}

func (ctl *MenuController) List() {
	// 参数对象
	var req dto.MenuQueryReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用查询方法
	list, err := services.Menu.GetList(req)
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

func (ctl *MenuController) Edit() {
	// 获取菜单列表
	menuTreeList, _ := services.Menu.GetTreeList()
	// 数据源转换
	menuList := services.Menu.MakeList(menuTreeList)
	// 记录ID
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.Menu{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 获取节点
		funcList := make([]models.Menu, 0)
		orm.NewOrm().QueryTable(new(models.Menu)).
			Filter("pid", id).
			Filter("type", 1).
			Filter("mark", 1).
			All(&funcList)
		sortList := make([]interface{}, 0)
		for _, v := range funcList {
			sortList = append(sortList, v.Sort)
		}

		// 渲染模板
		ctl.Data["info"] = info
		ctl.Data["funcList"] = sortList
	} else {
		// 添加
		pid, _ := ctl.GetInt("pid", 0)
		var info models.Menu
		info.Pid = pid
		info.Status = 1
		info.Target = 1
		ctl.Data["info"] = info
		ctl.Data["funcList"] = make([]interface{}, 0)
	}
	// 渲染模板
	ctl.Data["menuList"] = menuList
	ctl.Data["typeList"] = constant.MENU_TYPE_LIST
	ctl.Layout = "public/form.html"
	ctl.TplName = "menu/edit.html"
}

func (ctl *MenuController) Add() {
	// 参数对象
	var req dto.MenuAddReq
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
	// 调用添加方法
	rows, err := services.Menu.Add(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 添加成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (ctl *MenuController) Update() {
	// 参数对象
	var req dto.MenuUpdateReq
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
	// 调用更新方法
	rows, err := services.Menu.Update(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 更新成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "更新成功",
	})
}

func (ctl *MenuController) Delete() {
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.Menu.Delete(ids)
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 删除成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

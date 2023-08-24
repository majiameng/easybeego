/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/app/services"
	"easybeego/utils"
	"easybeego/utils/common"
	"github.com/gookit/validate"
)

var MemberLevel = new(MemberLevelController)

type MemberLevelController struct {
	BaseController
}

func (ctl *MemberLevelController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "member_level/index.html"
}

func (ctl *MemberLevelController) List() {
	// 查询对象
	var req dto.MemberLevelPageReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用查询方法
	list, count, err := services.MemberLevel.GetList(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "查询成功",
		Count: count,
	})
}

func (ctl *MemberLevelController) Edit() {
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.MemberLevel{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 渲染模板
		ctl.Data["info"] = info
	}
	ctl.Layout = "public/form.html"
	ctl.TplName = "member_level/edit.html"
}

func (ctl *MemberLevelController) Add() {
	// 添加对象
	var req dto.MemberLevelAddReq
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
	rows, err := services.MemberLevel.Add(req, utils.Uid(ctl.Ctx))
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

func (ctl *MemberLevelController) Update() {
	// 更新对象
	var req dto.MemberLevelUpdateReq
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
	rows, err := services.MemberLevel.Update(req, utils.Uid(ctl.Ctx))
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

func (ctl *MemberLevelController) Delete() {
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.MemberLevel.Delete(ids)
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

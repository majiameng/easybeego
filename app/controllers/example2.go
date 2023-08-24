/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 演示二管理-控制器
 * @author Tinymeng
 * @since 2022-05-13
 * @File : example2
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

var Example2 = new(Example2Controller)

type Example2Controller struct {
	BaseController
}

func (ctl *Example2Controller) Index() {
	// 模板渲染
	ctl.Layout = "public/layout.html"
	ctl.TplName = "example2/index.html"
}

func (ctl *Example2Controller) List() {
	// 参数绑定
	var req dto.Example2PageReq
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 调用获取列表方法
	lists, count, err := services.Example2.GetList(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code:  0,
		Data:  lists,
		Msg:   "操作成功",
		Count: count,
	})
}

func (ctl *Example2Controller) Edit() {
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.Example2{Id: id}
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
	ctl.TplName = "example2/edit.html"
}

func (ctl *Example2Controller) Add() {
	// 参数绑定
	var req dto.Example2AddReq
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
	rows, err := services.Example2.Add(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "添加成功",
	})
}

func (ctl *Example2Controller) Update() {
	// 参数绑定
	var req dto.Example2UpdateReq
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用更新方法
	rows, err := services.Example2.Update(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
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

func (ctl *Example2Controller) Delete() {
	// 记录ID
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}

	// 调用删除方法
	rows, err := services.Example2.Delete(ids)
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "删除成功",
	})
}

func (ctl *Example2Controller) Status() {
	// 参数绑定
	var req dto.Example2StatusReq
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 调用设置状态方法
	rows, err := services.Example2.Status(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}

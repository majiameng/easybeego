/**
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/app/services"
	"easybeego/utils"
	"easybeego/utils/common"
	"fmt"
	"github.com/gookit/validate"
)

// 实体对象
var Level = new(LevelController)

// 控制器结构体
type LevelController struct {
	BaseController
}

// 列表页视图
func (ctl *LevelController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "level/index.html"
}

// 查询数据列表
func (ctl *LevelController) List() {
	// 请求参数
	var req dto.LevelPageReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用获取列表函数
	list, count, err := services.Level.GetList(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果集
	ctl.JSON(common.JsonResult{
		Code:  0,
		Data:  list,
		Msg:   "查询成功",
		Count: count,
	})
}

// 编辑表单视图
func (ctl *LevelController) Edit() {
	// 查询记录
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		info := &models.Level{Id: id}
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
	// 渲染模板
	ctl.Layout = "public/form.html"
	ctl.TplName = "level/edit.html"
}

// 添加职级
func (ctl *LevelController) Add() {
	var req dto.LevelAddReq
	// 请求验证
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 表单验证
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}
	// 调用添加方法
	id, err := services.Level.Add(req, utils.Uid(ctl.Ctx))
	if err != nil || id == 0 {
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

// 更新职级
func (ctl *LevelController) Update() {
	// 参数验证
	var req dto.LevelUpdateReq
	// 请求验证
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 表单验证
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}
	// 调用更新方法
	rows, err := services.Level.Update(req, utils.Uid(ctl.Ctx))
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

// 删除职级
func (ctl *LevelController) Delete() {
	// 记录ID
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.Level.Delete(ids)
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

// 设置状态
func (ctl *LevelController) Status() {
	// 参数验证
	var req dto.LevelStatusReq
	// 请求验证
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 表单验证
	v := validate.Struct(req)
	if !v.Validate() {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  v.Errors.One(),
		})
	}
	// 调用更新方法
	rows, err := services.Level.Status(req, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 设置成功
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "设置成功",
	})
}

func (ctl *LevelController) ImportExcel() {
	// 调用上传方法
	result, err := services.Upload.UploadImage(ctl.Ctx)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用文件上传导入方法
	count, err := services.Level.ImportExcel(result.FileUrl, utils.Uid(ctl.Ctx))
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  fmt.Sprintf("本次共导入【%d】条数据", count),
	})
}

func (ctl *LevelController) ExportExcel() {
	// 请求参数
	var req dto.LevelPageReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用导出Excel方法
	fileURL, err := services.Level.GetExcelList(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "导出Excel失败",
		})
	}
	// 返回结果集
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "导出成功",
		Data: fileURL,
	})
}

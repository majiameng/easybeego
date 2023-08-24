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
	"easybeego/conf"
	"easybeego/utils"
	"easybeego/utils/common"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gookit/validate"
	"strings"
)

var Ad = new(AdController)

type AdController struct {
	BaseController
}

func (ctl *AdController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "ad/index.html"
}

func (ctl *AdController) List() {
	// 查询对象
	var req dto.AdPageReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用查询方法
	lists, count, err := services.Ad.GetList(req)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code:  0,
		Msg:   "查询成功",
		Data:  lists,
		Count: count,
	})
}

func (ctl *AdController) Edit() {
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.Ad{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		// 广告图片
		if info.Cover != "" {
			info.Cover = utils.GetImageUrl(info.Cover)
		}

		// 富文本图片替换处理
		if info.Content != "" {
			info.Content = strings.ReplaceAll(info.Content, "[IMG_URL]", conf.CONFIG.EGAdmin.Image)
		}

		// 渲染模板
		ctl.Data["info"] = info
	}
	// 广告位列表
	list := make([]models.AdSort, 0)
	orm.NewOrm().QueryTable(new(models.AdSort)).Filter("mark", 1).All(&list)
	adSortList := make(map[int]string, 0)
	for _, v := range list {
		adSortList[v.Id] = v.Description
	}
	// 投放平台
	ctl.Data["typeList"] = constant.AD_TYPE_LIST
	ctl.Data["adSortList"] = adSortList
	ctl.Layout = "public/form.html"
	ctl.TplName = "ad/edit.html"
}

func (ctl *AdController) Add() {
	// 添加对象
	var req dto.AdAddReq
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
	rows, err := services.Ad.Add(req, utils.Uid(ctl.Ctx))
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

func (ctl *AdController) Update() {
	// 更新对象
	var req dto.AdUpdateReq
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
	rows, err := services.Ad.Update(req, utils.Uid(ctl.Ctx))
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

func (ctl *AdController) Delete() {
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.Ad.Delete(ids)
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

func (ctl *AdController) Status() {
	// 设置状态对象
	var req dto.AdStatusReq
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
	// 调用设置状态方法
	rows, err := services.Ad.Status(req, utils.Uid(ctl.Ctx))
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

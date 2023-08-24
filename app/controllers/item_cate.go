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
	"github.com/beego/beego/v2/client/orm"
	"github.com/gookit/validate"
)

var ItemCate = new(ItemCateController)

type ItemCateController struct {
	BaseController
}

func (ctl *ItemCateController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "item_cate/index.html"
}

func (ctl *ItemCateController) List() {
	// 查询对象
	var req dto.ItemCateQueryReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用查询列表接口
	lists := services.ItemCate.GetList(req)
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "查询成功",
		Data: lists,
	})
}

func (ctl *ItemCateController) Edit() {
	// 记录ID
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.ItemCate{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 封面
		if info.IsCover == 1 && info.Cover != "" {
			info.Cover = utils.GetImageUrl(info.Cover)
		}
		// 渲染模板
		ctl.Data["info"] = info
	} else {
		// 渲染模板
		ctl.Data["info"] = &models.ItemCate{}
	}
	// 站点列表
	result := make([]models.Item, 0)
	orm.NewOrm().QueryTable(new(models.Item)).Filter("mark", 1).All(&result)
	var itemList = map[int]string{}
	for _, v := range result {
		itemList[v.Id] = v.Name
	}
	ctl.Data["itemList"] = itemList
	ctl.Layout = "public/form.html"
	ctl.TplName = "item_cate/edit.html"
}

func (ctl *ItemCateController) Add() {
	// 添加对象
	var req dto.ItemCateAddReq
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
	rows, err := services.ItemCate.Add(req, utils.Uid(ctl.Ctx))
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

func (ctl *ItemCateController) Update() {
	// 更新对象
	var req dto.ItemCateUpdateReq
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
	rows, err := services.ItemCate.Update(req, utils.Uid(ctl.Ctx))
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

func (ctl *ItemCateController) Delete() {
	// 记录ID
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.ItemCate.Delete(ids)
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

func (ctl *ItemCateController) GetCateTreeList() {
	itemId, _ := ctl.GetInt("itemId", 0)
	list, err := services.ItemCate.GetCateTreeList(itemId, 0)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}
	// 数据源转换
	result := services.ItemCate.MakeList(list)
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "操作成功",
		Data: result,
	})
}

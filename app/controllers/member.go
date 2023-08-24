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

var Member = new(MemberController)

type MemberController struct {
	BaseController
}

func (ctl *MemberController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "member/index.html"
}

func (ctl *MemberController) List() {
	// 查询对象
	var req dto.MemberPageReq
	// 参数绑定
	if err := ctl.ParseForm(&req); err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 调用查询方法
	list, count, err := services.Member.GetList(req)
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

func (ctl *MemberController) Edit() {
	// 会员等级
	list := make([]models.MemberLevel, 0)
	orm.NewOrm().QueryTable(new(models.MemberLevel)).Filter("mark", 1).All(&list)
	memberLevelList := make(map[int]string, 0)
	for _, v := range list {
		memberLevelList[v.Id] = v.Name
	}
	// 记录ID
	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.Member{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}
		// 头像
		if info.Avatar != "" {
			info.Avatar = utils.GetImageUrl(info.Avatar)
		}
		// 渲染模板
		ctl.Data["info"] = info
	}
	ctl.Data["memberLevelList"] = memberLevelList
	ctl.Data["deviceList"] = constant.MEMBER_DEVICE_LIST
	ctl.Data["sourceList"] = constant.MEMBER_SOURCE_LIST
	ctl.Layout = "public/form.html"
	ctl.TplName = "member/edit.html"
}

func (ctl *MemberController) Add() {
	// 添加对象
	var req dto.MemberAddReq
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
	rows, err := services.Member.Add(req, utils.Uid(ctl.Ctx))
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

func (ctl *MemberController) Update() {
	// 更新对象
	var req dto.MemberUpdateReq
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
	rows, err := services.Member.Update(req, utils.Uid(ctl.Ctx))
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

func (ctl *MemberController) Delete() {
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.Member.Delete(ids)
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

func (ctl *MemberController) Status() {
	// 参数对象
	var req dto.MemberStatusReq
	// 绑定参数
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
	rows, err := services.Member.Status(req, utils.Uid(ctl.Ctx))
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

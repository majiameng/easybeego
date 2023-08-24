/**
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/constant"
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/app/services"
	"easybeego/app/vo"
	"easybeego/utils"
	"easybeego/utils/common"
	"github.com/beego/beego/v2/client/orm"
	"github.com/gookit/validate"
)

var User = new(UserController)

type UserController struct {
	BaseController
}

func (ctl *UserController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "user/index.html"
}

func (ctl *UserController) List() {
	// 参数对象
	var req dto.UserPageReq
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
	// 调用查询方法
	list, count, err := services.User.GetList(req)
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

func (ctl *UserController) Edit() {
	// 获取职级
	levelAll := make([]models.Level, 0)
	orm.NewOrm().QueryTable(new(models.Level)).Filter("status", 1).Filter("mark", 1).All(&levelAll)
	levelList := make(map[int]string, 0)
	for _, v := range levelAll {
		levelList[v.Id] = v.Name
	}
	// 获取岗位
	positionAll := make([]models.Position, 0)
	orm.NewOrm().QueryTable(new(models.Position)).Filter("status", 1).Filter("mark", 1).All(&positionAll)
	positionList := make(map[int]string, 0)
	for _, v := range positionAll {
		positionList[v.Id] = v.Name
	}
	// 获取部门列表
	deptData, _ := services.Dept.GetDeptTreeList()
	deptList := services.Dept.MakeList(deptData)
	// 获取角色
	roleData := make([]models.Role, 0)
	orm.NewOrm().QueryTable(new(models.Role)).Filter("status", 1).Filter("mark", 1).All(&roleData)
	roleList := make(map[int]string)
	for _, v := range roleData {
		roleList[v.Id] = v.Name
	}

	id, _ := ctl.GetInt("id", 0)
	if id > 0 {
		// 编辑
		info := &models.User{Id: id}
		err := info.Get()
		if err != nil {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
		}

		var userInfo = vo.UserInfoVo{}
		userInfo.User = *info
		// 头像
		userInfo.Avatar = utils.GetImageUrl(info.Avatar)

		// 角色ID
		var userRoleList []models.UserRole
		orm.NewOrm().QueryTable(new(models.UserRole)).Filter("user_id", info.Id).All(&userRoleList)
		roleIds := make([]interface{}, 0)
		for _, v := range userRoleList {
			roleIds = append(roleIds, v.RoleId)
		}
		userInfo.RoleIds = roleIds

		// 渲染模板
		ctl.Data["info"] = userInfo
	}
	ctl.Data["genderList"] = constant.GENDER_LIST
	ctl.Data["levelList"] = levelList
	ctl.Data["positionList"] = positionList
	ctl.Data["deptList"] = deptList
	ctl.Data["roleList"] = roleList
	ctl.Layout = "public/form.html"
	ctl.TplName = "user/edit.html"
}

func (ctl *UserController) Add() {
	// 参数对象
	var req dto.UserAddReq
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
	rows, err := services.User.Add(req, utils.Uid(ctl.Ctx))
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

func (ctl *UserController) Update() {
	// 参数对象
	var req dto.UserUpdateReq
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
	rows, err := services.User.Update(req, utils.Uid(ctl.Ctx))
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

func (ctl *UserController) Delete() {
	ids := ctl.GetString("id")
	if ids == "" {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  "记录ID不能为空",
		})
	}
	// 调用删除方法
	rows, err := services.User.Delete(ids)
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

func (ctl *UserController) Status() {
	// 参数对象
	var req dto.UserStatusReq
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
	rows, err := services.User.Status(req, utils.Uid(ctl.Ctx))
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

func (ctl *UserController) ResetPwd() {
	// 重置密码对象
	var req dto.UserResetPwdReq
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
	// 调用重置密码方法
	rows, err := services.User.ResetPwd(req.Id, utils.Uid(ctl.Ctx))
	if err != nil || rows == 0 {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}

	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "重置密码成功",
	})
}

/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 分页查询条件
type MemberPageReq struct {
	Username string `form:"username"` // 用户名
	Gender   int    `form:"gender"`   // 性别（1男 2女 3未知）
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加会员
type MemberAddReq struct {
	Username     string `form:"username,unique" validate:"required"` // 用户名
	Password     string `form:"password"`                            // 登录密码
	MemberLevel  int    `form:"memberLevel" validate:"int"`          // 会员等级
	Realname     string `form:"realname" validate:"required"`        // 真实姓名
	Nickname     string `form:"nickname" validate:"required"`        // 用户昵称
	Gender       int    `form:"gender" validate:"int"`               // 性别（1男 2女 3未知）
	Avatar       string `form:"avatar" validate:"required"`          // 用户头像
	Birthday     string `form:"birthday" validate:"required"`        // 出生日期
	ProvinceCode string `form:"provinceCode" validate:"required"`    // 省份编号
	CityCode     string `form:"cityCode" validate:"required"`        // 市区编号
	DistrictCode string `form:"districtCode" validate:"required"`    // 区县编号
	Address      string `form:"address" validate:"required"`         // 详细地址
	Intro        string `form:"intro"`                               // 个人简介
	Signature    string `form:"signature"`                           // 个性签名
	Device       int    `form:"device" validate:"int"`               // 设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加
	Source       int    `form:"source" validate:"int"`               // 来源：1、APP注册；2、后台添加；
	Status       int    `form:"status" validate:"int"`               // 是否启用 1、启用  2、停用
}

// 添加会员表单验证
func (v MemberAddReq) Messages() map[string]string {
	return validate.MS{
		"Username.required":     "用户名不能为空.",
		"MemberLevel.int":       "请选择会员等级.",
		"Realname.required":     "真实姓名不能为空.",
		"Nickname.required":     "昵称不能为空.",
		"Gender.int":            "请选择性别.",
		"Avatar.required":       "请上传头像.",
		"Birthday.required":     "请选择出生日期.",
		"ProvinceCode.required": "请选择省份.",
		"CityCode.required":     "请选择城市.",
		"DistrictCode.required": "请选择县区.",
		"Device.int":            "请选择设备类型.",
		"Source.int":            "请选择注册来源.",
		"Status.int":            "请选择会员状态.",
	}
}

// 更新会员
type MemberUpdateReq struct {
	Id           int    `form:"id" validate:"int"`
	Username     string `form:"username,unique" validate:"required"` // 用户名
	Password     string `form:"password"`                            // 登录密码
	MemberLevel  int    `form:"memberLevel" validate:"int"`          // 会员等级
	Realname     string `form:"realname" validate:"required"`        // 真实姓名
	Nickname     string `form:"nickname" validate:"required"`        // 用户昵称
	Gender       int    `form:"gender" validate:"int"`               // 性别（1男 2女 3未知）
	Avatar       string `form:"avatar" validate:"required"`          // 用户头像
	Birthday     string `form:"birthday" validate:"required"`        // 出生日期
	ProvinceCode string `form:"provinceCode" validate:"required"`    // 省份编号
	CityCode     string `form:"cityCode" validate:"required"`        // 市区编号
	DistrictCode string `form:"districtCode" validate:"required"`    // 区县编号
	Address      string `form:"address" validate:"required"`         // 详细地址
	Intro        string `form:"intro"`                               // 个人简介
	Signature    string `form:"signature"`                           // 个性签名
	Device       int    `form:"device" validate:"int"`               // 设备类型：1苹果 2安卓 3WAP站 4PC站 5后台添加
	Source       int    `form:"source" validate:"int"`               // 来源：1、APP注册；2、后台添加；
	Status       int    `form:"status" validate:"int"`               // 是否启用 1、启用  2、停用
}

// 更新会员表单验证
func (v MemberUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":                "会员ID不能为空.",
		"Username.required":     "用户名不能为空.",
		"MemberLevel.int":       "请选择会员等级.",
		"Realname.required":     "真实姓名不能为空.",
		"Nickname.required":     "昵称不能为空.",
		"Gender.int":            "请选择性别.",
		"Avatar.required":       "请上传头像.",
		"Birthday.required":     "请选择出生日期.",
		"ProvinceCode.required": "请选择省份.",
		"CityCode.required":     "请选择城市.",
		"DistrictCode.required": "请选择县区.",
		"Device.int":            "请选择设备类型.",
		"Source.int":            "请选择注册来源.",
		"Status.int":            "请选择会员状态.",
	}
}

// 设置状态
type MemberStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

func (v MemberStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "会员ID不能为空.",
		"Status.int": "请选择会员状态.",
	}
}

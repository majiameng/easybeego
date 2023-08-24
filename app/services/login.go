/**
 * @author: Tinymeng <666@majiameng.com>
 */

package services

import (
	"easybeego/app/models"
	"easybeego/conf"
	"easybeego/utils"
	"easybeego/utils/gstr"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"time"
)

var Login = new(loginService)

type loginService struct{}

// 系统登录
func (s *loginService) UserLogin(username, password string, ctx *context.Context) (*models.User, error) {
	// 查询用户
	var user models.User
	err := orm.NewOrm().QueryTable(new(models.User)).Filter("username", username).Filter("mark", 1).One(&user)
	if err != nil {
		return nil, errors.New("用户名或者密码不正确")
	}
	// 密码校验
	pwd, _ := utils.Md5(password + user.Username)
	if user.Password != pwd {
		return nil, errors.New("密码不正确")
	}
	// 判断当前用户状态
	if user.Status != 1 {
		return nil, errors.New("您的账号已被禁用,请联系管理员")
	}

	// 更新登录时间、登录IP
	o := orm.NewOrm()
	entity := models.User{Id: user.Id}
	entity.LoginTime = time.Now()
	entity.LoginIp = ctx.Input.IP()
	entity.UpdateTime = time.Now()
	o.Update(&entity, "LoginTime", "LoginIp", "UpdateTime")

	// 结果
	return &user, nil
}

// 获取个人信息
func (s *loginService) GetProfile(userId int) *models.User {
	user := &models.User{Id: userId}
	err := user.Get()
	if err != nil {
		return nil
	}
	// 头像
	if user.Avatar != "" && !gstr.Contains(user.Avatar, conf.CONFIG.EGAdmin.Image) {
		user.Avatar = utils.GetImageUrl(user.Avatar)
	}
	return user
}

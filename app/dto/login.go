/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 系统登录
type LoginReq struct {
	UserName string `form:"username" validate:"required"`
	Password string `form:"password" validate:"required"`
	Captcha  string `form:"captcha" validate:"required"`
	IdKey    string `form:"idKey" validate:"required"`
}

// 登录参数校验
func (v LoginReq) Messages() map[string]string {
	return validate.MS{
		"UserName.required": "登录用户名不能为空.",
		"Password.required": "登录密码不能为空.",
		"Captcha.required":  "登录验证码不能为空.",
		"IdKey.required":    "验证码KEY不能为空.",
	}
}

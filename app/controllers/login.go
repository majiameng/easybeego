/**
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/dto"
	"easybeego/app/services"
	"easybeego/conf"
	"easybeego/utils/common"
	"github.com/gookit/validate"
	"github.com/mojocn/base64Captcha"
)

var Login = new(LoginController)

type LoginController struct {
	BaseController
}

// 系统登录
func (ctl *LoginController) Login() {
	if ctl.Ctx.Input.IsPost() {
		// 登录对象
		var req dto.LoginReq
		// 参数绑定
		if err := ctl.ParseForm(&req); err != nil {
			// 返回错误信息
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  "登录错误",
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

		// 校验验证码
		//verifyRes := base64Captcha.VerifyCaptcha(req.IdKey, req.Captcha)
		//if !verifyRes {
		//	ctl.JSON(common.JsonResult{
		//		Code: -1,
		//		Msg:  "验证码不正确",
		//	})
		//	return
		//}

		// 系统登录
		user, err := services.Login.UserLogin(req.UserName, req.Password, ctl.Ctx)
		if err != nil {
			// 登录错误
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  err.Error(),
			})
			return
		}

		// 保存Session
		// 设置session 记住用户登录状态
		ctl.SetSession(conf.USER_ID, user.Id)
		//userId := ctl.GetSession("userId")
		//ctl.DelSession("userId")
		// 登录成功
		ctl.JSON(common.JsonResult{
			Code: 0,
			Msg:  "登录成功",
		})
	}
	// 渲染模板
	ctl.TplName = "login.html"
}

// 验证码
func (ctl *LoginController) Captcha() {
	// 验证码参数配置：字符,公式,验证码配置
	var configC = base64Captcha.ConfigCharacter{
		Height: 60,
		Width:  240,
		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
		Mode:               base64Captcha.CaptchaModeAlphabet,
		ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
		ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         6,
	}
	///create a characters captcha.
	idKeyC, capC := base64Captcha.GenerateCaptcha("", configC)
	//以base64编码
	base64stringC := base64Captcha.CaptchaWriteToBase64Encoding(capC)

	// 返回结果集
	ctl.JSON(common.CaptchaRes{
		Code:  0,
		IdKey: idKeyC,
		Data:  base64stringC,
		Msg:   "操作成功",
	})
}

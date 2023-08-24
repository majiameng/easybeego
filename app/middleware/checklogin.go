/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 登录验证中间件
 * @author 半城风雨
 * @since 2021/8/20
 * @File : checkauth
 */
package middleware

import (
	"easybeego/conf"
	"easybeego/utils"
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/adapter/context"
	"net/http"
	"strings"
)

func CheckLogin() {

	//登录认证中间件过滤器
	var login = func(ctx *context.Context) {
		fmt.Println("登录验证中间件")
		// 放行设置
		urlItem := []string{"/captcha", "/login", "/user/index"}
		if !utils.InStringArray(ctx.Request.URL.Path, urlItem) && !strings.Contains(ctx.Request.URL.Path, "resource") {
			if !IsLogin(ctx) {
				// 跳转登录页,方式：301(永久移动),308(永久重定向),307(临时重定向)
				ctx.Redirect(http.StatusFound, "/login")
				return
			}
		}
	}

	// 登录过滤器
	beego.InsertFilter("/*", beego.BeforeRouter, login)
}

func IsLogin(ctx *context.Context) bool {
	userId := ctx.Input.Session(conf.USER_ID)
	return userId != nil
}

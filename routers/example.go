/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 演示一-路由
 * @author Tinymeng
 * @since 2022-05-13
 * @File : example
 */
package routers

import (
	"easybeego/app/controllers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	fmt.Println("模块路由初始化")

	/* 演示一 */
	beego.Router("/example/index", &controllers.ExampleController{}, "get:Index")
	beego.Router("/example/list", &controllers.ExampleController{}, "post:List")
	beego.Router("/example/edit", &controllers.ExampleController{}, "get:Edit")
	beego.Router("/example/add", &controllers.ExampleController{}, "post:Add")
	beego.Router("/example/update", &controllers.ExampleController{}, "post:Update")
	beego.Router("/example/delete", &controllers.ExampleController{}, "post:Delete")
	beego.Router("/example/setStatus", &controllers.ExampleController{}, "post:Status")

	beego.Router("/example/setStatus", &controllers.ExampleController{}, "post:Status")
	beego.Router("/example/setIsVip", &controllers.ExampleController{}, "post:IsVip")
}

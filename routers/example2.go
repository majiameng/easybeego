/**
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 演示二-路由
 * @author 半城风雨
 * @since 2022-05-13
 * @File : example2
 */
package routers

import (
	"easybeego/app/controllers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	fmt.Println("模块路由初始化")

	/* 演示二 */
	beego.Router("/example2/index", &controllers.Example2Controller{}, "get:Index")
	beego.Router("/example2/list", &controllers.Example2Controller{}, "post:List")
	beego.Router("/example2/edit", &controllers.Example2Controller{}, "get:Edit")
	beego.Router("/example2/add", &controllers.Example2Controller{}, "post:Add")
	beego.Router("/example2/update", &controllers.Example2Controller{}, "post:Update")
	beego.Router("/example2/delete", &controllers.Example2Controller{}, "post:Delete")
	beego.Router("/example2/setStatus", &controllers.Example2Controller{}, "post:Status")

	beego.Router("/example2/setStatus", &controllers.Example2Controller{}, "post:Status")
}

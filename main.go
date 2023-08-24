/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package main

import (
	_ "easybeego/boot/config"
	_ "easybeego/boot/mysql"
	_ "easybeego/boot/session"
	_ "easybeego/routers"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	orm.Debug = true
	// 启动应用
	beego.Run()
}

/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package mysql

import (
	"easybeego/conf"
	"fmt"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 注册MySQL
func init() {
	// 注册数据库驱动
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		beego.Error("mysql register driver error:", err)
	}

	//dataSource := "root:root@tcp(127.0.0.1:3306)/demo"
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conf.CONFIG.Mysql.Username,
		conf.CONFIG.Mysql.Password,
		conf.CONFIG.Mysql.Host,
		conf.CONFIG.Mysql.Port,
		conf.CONFIG.Mysql.Database,
	)

	// 注册数据库
	err = orm.RegisterDataBase("default", "mysql", dataSource)
	if err != nil {
		beego.Error("mysql register database error:", err)
	}
}

/**
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	beego "github.com/beego/beego/v2/adapter"
)

// 基类结构体
type BaseController struct {
	beego.Controller
}

type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：0成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
	Count int64       `json:"count"` // 记录总数
}

func (this *BaseController) JSON(obj interface{}) {
	this.Data["json"] = obj
	//对json进行序列化输出
	this.ServeJSON()
	this.StopRun()
}

//func (ctl *BaseController) Html(params ...string) {
//	ctl.Ctx.WriteString(params[0])
//}
//
//func (ctl *BaseController) Render(params ...string) {
//}

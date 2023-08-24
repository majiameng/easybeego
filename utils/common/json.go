/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 公共函数库
 * @author 半城风雨
 * @since 2021/3/2
 * @File : common
 */
package common

type BunissType int

// 业务类型
const (
	BOther BunissType = 0 //0其它
	BAdd   BunissType = 1 //1新增
	BEdit  BunissType = 2 //2修改
	BDel   BunissType = 3 //3删除
)

type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：0成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
	Count int64       `json:"count"` // 记录总数
}

type CaptchaRes struct {
	Code  int         `json:"code"`  //响应编码 0 成功 500 错误 403 无权限
	Msg   string      `json:"msg"`   //消息
	Data  interface{} `json:"data"`  //数据内容
	IdKey string      `json:"idkey"` //验证码ID
}

type JsonEditResult struct {
	Error int    `json:"error"` // 错误编码
	Url   string `json:"url"`   // 图片地址
}

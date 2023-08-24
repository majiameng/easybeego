/**
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

// 分页查询
type Example2PageReq struct {
	Name string `form:"name"` // 演示名称

	Status int `form:"status"` // 状态：1正常 2停用

	Page  int `form:"page"`  // 页码
	Limit int `form:"limit"` // 每页数
}

// 添加演示二
type Example2AddReq struct {
	Name   string `form:"name" validate:"required"` // 演示名称
	Status int    `form:"status" validate:"int"`    // 状态：1正常 2停用
	Sort   int    `form:"sort" validate:"int"`      // 排序号
}

// 编辑演示二
type Example2UpdateReq struct {
	Id string `form:"id" validate:"required"`

	Name string `form:"name" validate:"required"` // 演示名称

	Status int `form:"status" validate:"int"` // 状态：1正常 2停用

	Sort int `form:"sort" validate:"int"` // 排序号

}

// 设置状态
type Example2StatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

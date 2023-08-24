/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package dto

import "github.com/gookit/validate"

// 列表查询条件
type CityQueryReq struct {
	Name string `form:"name"` // 城市名称
	Pid  int    `form:"pid"`  // 上级ID
}

// 添加城市
type CityAddReq struct {
	Pid      int    `form:"pid"`                      // 父级编号
	Level    int    `form:"level" validate:"int"`     // 城市级别：1省 2市 3区
	Name     string `form:"name" validate:"required"` // 城市名称
	Citycode string `form:"citycode"`                 // 城市编号（区号）
	PAdcode  string `form:"pAdcode"`                  // 父级地理编号
	Adcode   string `form:"adcode"`                   // 地理编号
	Lng      string `form:"lng"`                      // 城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7
	Lat      string `form:"lat"`                      // 城市坐标中心点纬度（* 1e6）
	Sort     int    `form:"sort" validate:"int"`      // 排序号
}

// 添加城市表单验证
func (v CityAddReq) Messages() map[string]string {
	return validate.MS{
		"Level.int":     "请选择城市级别.",
		"Name.required": "城市名称不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

// 编辑城市
type CityUpdateReq struct {
	Id       int    `form:"id" validate:"int"`        // 主键ID
	Pid      int    `form:"pid"`                      // 父级编号
	Level    int    `form:"level" validate:"int"`     // 城市级别：1省 2市 3区
	Name     string `form:"name" validate:"required"` // 城市名称
	Citycode string `form:"citycode"`                 // 城市编号（区号）
	PAdcode  string `form:"pAdcode"`                  // 父级地理编号
	Adcode   string `form:"adcode"`                   // 地理编号
	Lng      string `form:"lng"`                      // 城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7
	Lat      string `form:"lat"`                      // 城市坐标中心点纬度（* 1e6）
	Sort     int    `form:"sort" validate:"int"`      // 排序号
}

// 编辑城市表单验证
func (v CityUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "城市ID不能为空.",
		"Level.int":     "请选择城市级别.",
		"Name.required": "城市名称不能为空.",
		"Sort.int":      "排序不能为空.",
	}
}

// 获取子级城市
type CityChildReq struct {
	CityCode string `form:"citycode" validate:"required"`
}

// 获取子级城市参数验证
func (v CityChildReq) Messages() map[string]string {
	return validate.MS{
		"CityCode.required": "城市编码不能为空.",
	}
}

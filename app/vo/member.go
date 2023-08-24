/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 会员信息Vo
type MemberInfoVo struct {
	models.Member
	GenderName string      `json:"GenderName"` // 性别
	DeviceName string      `json:"DeviceName"` // 设备类型
	SourceName string      `json:"SourceName"` // 会员来源
	City       interface{} `json:"City"`       // 省市区
	CityName   string      `json:"CityName"`   // 城市名称
}

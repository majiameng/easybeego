/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 配置数据列表
type ConfigDataVo struct {
	models.ConfigData
	TypeName string `json:"TypeName"`
}

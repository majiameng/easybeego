/**
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 友链信息
type LinkInfoVo struct {
	models.Link
	TypeName     string `json:"TypeName"`     // 友链类型
	FormName     string `json:"FormName"`     // 友链形式
	PlatformName string `json:"PlatformName"` // 投放平台
}

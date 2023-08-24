/**
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 站点信息Vo
type ItemInfoVo struct {
	models.Item
	TypeName string `json:"TypeName"` // 站点类型
}

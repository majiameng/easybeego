/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 广告信息Vo
type AdInfoVo struct {
	models.Ad
	TypeName   string `json:"TypeName"`   // 广告类型
	AdSortDesc string `json:"AdSortDesc"` // 广告位描述
}

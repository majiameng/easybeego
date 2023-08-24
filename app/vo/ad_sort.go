/**
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 广告位信息
type AdSortInfoVo struct {
	models.AdSort
	ItemName     string `json:"ItemName"`     // 站点名称
	CateName     string `json:"CateName"`     // 栏目名称
	PlatformName string `json:"PlatformName"` // 所属平台
}

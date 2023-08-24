/**
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 通知公告Vo
type NoticeInfoVo struct {
	models.Notice
	SourceName string `json:"SourceName"` // 通知来源
}

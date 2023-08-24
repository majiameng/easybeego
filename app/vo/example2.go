/**
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 演示二Vo
 * @author 半城风雨
 * @since 2022-05-13
 * @File : example2
 */
package vo

import "easybeego/app/models"

// 演示二信息Vo
type Example2InfoVo struct {
	models.Example2

	StatusName int `json:"statusName"` // 状态名称
}

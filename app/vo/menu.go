/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 菜单Vo
type MenuTreeNode struct {
	models.Menu
	Children []*MenuTreeNode `json:"children"` // 子菜单
}

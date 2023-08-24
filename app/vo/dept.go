/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 部门树结构
type DeptTreeNode struct {
	models.Dept
	Children []*DeptTreeNode `json:"children"` // 子栏目
}

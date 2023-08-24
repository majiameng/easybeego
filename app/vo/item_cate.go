/**
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

import "easybeego/app/models"

// 栏目信息
type ItemCateInfoVo struct {
	models.ItemCate
	ItemName string `json:"ItemName"` // 栏目名称
}

// 栏目树结构
type CateTreeNode struct {
	models.ItemCate
	Children []*CateTreeNode `json:"children"` // 子栏目
}

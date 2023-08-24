/**
 * @author: Tinymeng <666@majiameng.com>
 */

package services

import (
	"easybeego/app/models"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

var UserRole = new(userRoleService)

type userRoleService struct{}

// 获取用户角色列表
func (s *userRoleService) GetUserRoleList(userId int) []models.Role {
	// 实例化对象
	list := make([]models.Role, 0)
	// 查询SQL语句
	sql := "SELECT r.* FROM sys_role AS r " +
		" INNER JOIN sys_user_role AS ur ON r.id=ur.role_id" +
		" WHERE ur.user_id=" + strconv.Itoa(userId) + " AND r.mark=1" +
		" ORDER BY r.sort asc"
	// 执行查询并转换对象
	orm.NewOrm().Raw(sql).QueryRows(&list)
	return list
}

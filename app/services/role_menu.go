/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package services

import (
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/app/vo"
	"easybeego/utils"
	"easybeego/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strings"
)

var RoleMenu = new(roleMenuService)

type roleMenuService struct{}

func (s *roleMenuService) GetRoleMenuList(roleId int) ([]vo.RoleMenuInfo, error) {
	// 获取全部菜单列表
	var menuList []models.Menu
	orm.NewOrm().QueryTable(new(models.Menu)).
		Filter("status", 1).
		Filter("mark", 1).
		OrderBy("sort").
		All(&menuList)
	if len(menuList) == 0 {
		return nil, errors.New("菜单列表不存在")
	}
	// 获取角色菜单权限列表
	var roleMenuList []models.RoleMenu
	orm.NewOrm().QueryTable(new(models.RoleMenu)).Filter("role_id", roleId).All(&roleMenuList)
	idList := make([]interface{}, 0)
	for _, v := range roleMenuList {
		idList = append(idList, v.MenuId)
	}

	// 对象处理
	var list []vo.RoleMenuInfo
	if len(menuList) > 0 {
		for _, m := range menuList {
			var info vo.RoleMenuInfo
			info.Id = m.Id
			info.Name = m.Name
			info.Open = true
			info.Pid = m.Pid
			// 节点选中值
			if utils.InArray(gconv.String(m.Id), idList) {
				info.Checked = true
			}
			list = append(list, info)
		}
	}
	return list, nil
}

func (s *roleMenuService) Save(req dto.RoleMenuSaveReq) error {
	if utils.AppDebug() {
		return errors.New("演示环境，暂无权限操作")
	}
	itemArr := strings.Split(req.MenuIds, ",")
	// 删除现有的角色权限数据
	orm.NewOrm().QueryTable(new(models.RoleMenu)).Filter("role_id", req.RoleId).Delete()
	// 遍历创建新角色权限数据
	for _, v := range itemArr {
		var entity models.RoleMenu
		entity.RoleId = req.RoleId
		entity.MenuId = gconv.Int(v)
		entity.Insert()
	}
	return nil
}

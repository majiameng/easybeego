/**
 * @author: Tinymeng <666@majiameng.com>
 */

package services

import (
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/utils"
	"easybeego/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"time"
)

var MemberLevel = new(memberLevelService)

type memberLevelService struct{}

func (s *memberLevelService) GetList(req dto.MemberLevelPageReq) ([]models.MemberLevel, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.MemberLevel)).Filter("mark", 1)
	// 会员等级名称查询
	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询总数
	count, err := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.MemberLevel, 0)
	query.All(&lists)
	// 返回结果
	return lists, count, err
}

func (s *memberLevelService) Add(req dto.MemberLevelAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.MemberLevel
	entity.Name = req.Name
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *memberLevelService) Update(req dto.MemberLevelUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.MemberLevel{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}
	// 设置参数
	entity.Name = req.Name
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

func (s *memberLevelService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.MemberLevel{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.MemberLevel{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

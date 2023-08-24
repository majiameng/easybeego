/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 演示二管理-服务类
 * @author 半城风雨
 * @since 2022-05-13
 * @File : example2
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
	"time"
)

// 中间件管理服务
var Example2 = new(example2Service)

type example2Service struct{}

func (s *example2Service) GetList(req dto.Example2PageReq) ([]vo.Example2InfoVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Example2)).Filter("mark", 1)

	// 演示名称

	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}

	// 状态：1正常 2停用

	if req.Status > 0 {
		query = query.Filter("status", req.Status)
	}

	// 排序
	query = query.OrderBy("id")
	// 查询总数
	count, _ := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.Example2, 0)
	// 对象转换
	query.All(&lists)

	// 数据处理
	var result []vo.Example2InfoVo
	for _, v := range lists {
		item := vo.Example2InfoVo{}
		item.Example2 = v

		result = append(result, item)
	}

	// 返回结果
	return result, count, nil
}

func (s *example2Service) Add(req dto.Example2AddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.Example2

	entity.Name = req.Name

	entity.Status = req.Status

	entity.Sort = req.Sort

	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *example2Service) Update(req dto.Example2UpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.Example2{Id: gconv.Int(req.Id)}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}

	entity.Name = req.Name

	entity.Status = req.Status

	entity.Sort = req.Sort

	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

// 删除
func (s *example2Service) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Example2{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.Example2{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *example2Service) Status(req dto.Example2StatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.Example2{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}

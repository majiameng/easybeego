/**
 * @author: Tinymeng <666@majiameng.com>
 */

package services

import (
	"easybeego/app/constant"
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/app/vo"
	"easybeego/utils"
	"easybeego/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"strings"
	"time"
)

var ConfigData = new(configDataService)

type configDataService struct{}

func (s *configDataService) GetList(req dto.ConfigDataPageReq) ([]vo.ConfigDataVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.ConfigData)).Filter("config_id", req.ConfigId).Filter("mark", 1)
	// 配置项名称查询
	if req.Title != "" {
		query = query.Filter("title__contains", req.Title)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询总数
	count, _ := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.ConfigData, 0)
	// 查询并转换对象
	query.All(&lists)

	// 数据源处理
	var result = make([]vo.ConfigDataVo, 0)
	for _, v := range lists {
		typeName, ok := constant.CONFIG_DATA_TYPE_LIST[v.Type]
		item := vo.ConfigDataVo{}
		item.ConfigData = v
		if ok {
			item.TypeName = typeName
		}
		result = append(result, item)
	}
	// 返回结果
	return result, count, nil
}

func (s *configDataService) Add(req dto.ConfigDataAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.ConfigData
	entity.Title = req.Title
	entity.Code = req.Code
	entity.Value = req.Value
	entity.Options = req.Options
	entity.ConfigId = req.ConfigId
	entity.Type = req.Type
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *configDataService) Update(req dto.ConfigDataUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.ConfigData{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}
	// 设置对象
	entity.Title = req.Title
	entity.Code = req.Code
	entity.Value = req.Value
	entity.Options = req.Options
	entity.ConfigId = req.ConfigId
	entity.Type = req.Type
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

func (s *configDataService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 分裂记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.ConfigData{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			id, _ := strconv.Atoi(v)
			entity := &models.ConfigData{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *configDataService) Status(req dto.ConfigDataStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.ConfigData{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}

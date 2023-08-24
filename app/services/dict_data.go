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
	"strconv"
	"strings"
	"time"
)

var DictData = new(dictDataService)

type dictDataService struct{}

func (s *dictDataService) GetList(req dto.DictDataPageReq) ([]models.DictData, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.DictData)).Filter("dict_id", req.DictId).Filter("mark", 1)
	// 字典项名称查询
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
	lists := make([]models.DictData, 0)
	// 查询并转换对象
	query.All(&lists)
	// 返回结果
	return lists, count, err
}

func (s *dictDataService) Add(req dto.DictDataAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.DictData
	entity.DictId = req.DictId
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1

	// 插入记录
	return entity.Insert()
}

func (s *dictDataService) Update(req dto.DictDataUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.DictData{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}
	// 设置对象
	entity.DictId = req.DictId
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

func (s *dictDataService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 分裂记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.DictData{Id: gconv.Int(ids)}
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
			entity := &models.DictData{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

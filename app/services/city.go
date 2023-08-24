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
	"time"
)

var City = new(cityService)

type cityService struct{}

func (s *cityService) GetList(req dto.CityQueryReq) ([]vo.CityInfoVo, error) {
	// 创建查询实例
	query := orm.NewOrm().QueryTable(new(models.City)).Filter("mark", 1)
	// 上级城市ID
	query = query.Filter("pid", req.Pid)
	// 城市名称
	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询数据
	var list []models.City
	_, err := query.All(&list)

	// 数据解析
	var result = make([]vo.CityInfoVo, 0)
	// 遍历数据
	for _, v := range list {
		item := vo.CityInfoVo{}
		item.City = v
		if v.Level < 3 {
			item.HaveChild = true
		} else {
			item.HaveChild = false
		}
		result = append(result, item)
	}
	return result, err
}

func (s *cityService) Add(req dto.CityAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.City
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.Level = req.Level
	entity.Citycode = req.Citycode
	entity.PAdcode = req.PAdcode
	entity.Adcode = req.Adcode
	entity.Lng = req.Lng
	entity.Lat = req.Lat
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入记录
	return entity.Insert()
}

func (s *cityService) Update(req dto.CityUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.City{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}
	// 设置对象属性
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.Level = req.Level
	entity.Citycode = req.Citycode
	entity.PAdcode = req.PAdcode
	entity.Adcode = req.Adcode
	entity.Lng = req.Lng
	entity.Lat = req.Lat
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

func (s *cityService) GetChilds(cityCode string) ([]models.City, error) {
	var info models.City
	o := orm.NewOrm()
	err := o.QueryTable(new(models.City)).Filter("citycode", cityCode).Filter("mark", 1).One(&info)
	if err != nil {
		return nil, errors.New("城市不能存在")
	}
	list := make([]models.City, 0)
	o.QueryTable(new(models.City)).Filter("pid", info.Id).Filter("mark", 1).All(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *cityService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.City{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.City{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *cityService) GetCityName(cityCode string, delimiter string) string {
	var info models.City
	err := orm.NewOrm().QueryTable(new(models.City)).Filter("citycode", cityCode).One(&info)
	if err != nil {
		return ""
	}
	// 城市ID
	cityId := info.Id
	// 声明数组
	list := make([]string, 0)
	for {
		if cityId <= 0 {
			// 退出
			break
		}
		// 业务处理
		info = models.City{}
		err2 := orm.NewOrm().QueryTable(new(models.City)).Filter("id", cityId).One(&info)
		if err2 != nil {
			break
		}
		// 上级栏目ID
		cityId = info.Pid
		// 加入数组
		list = append(list, info.Name)
	}
	// 结果数据处理
	if len(list) > 0 {
		// 数组反转
		utils.Reverse(&list)
		// 拼接字符串
		return strings.Join(list, delimiter)
	}
	return ""
}

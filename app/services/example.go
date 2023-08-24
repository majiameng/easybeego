/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

/**
 * 演示一管理-服务类
 * @author 半城风雨
 * @since 2022-05-13
 * @File : example
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
var Example = new(exampleService)

type exampleService struct{}

func (s *exampleService) GetList(req dto.ExamplePageReq) ([]vo.ExampleInfoVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Example)).Filter("mark", 1)

	// 测试名称

	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}

	// 状态：1正常 2停用

	if req.Status > 0 {
		query = query.Filter("status", req.Status)
	}

	// 类型：1京东 2淘宝 3拼多多 4唯品会

	if req.Type > 0 {
		query = query.Filter("type", req.Type)
	}

	// 是否VIP：1是 2否

	if req.IsVip > 0 {
		query = query.Filter("is_vip", req.IsVip)
	}

	// 排序
	query = query.OrderBy("id")
	// 查询总数
	count, _ := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.Example, 0)
	// 对象转换
	query.All(&lists)

	// 数据处理
	var result []vo.ExampleInfoVo
	for _, v := range lists {
		item := vo.ExampleInfoVo{}
		item.Example = v

		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}

		result = append(result, item)
	}

	// 返回结果
	return result, count, nil
}

func (s *exampleService) Add(req dto.ExampleAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.Example

	entity.Name = req.Name
	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "example")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.Content = req.Content

	entity.Status = req.Status

	entity.Type = req.Type

	entity.IsVip = req.IsVip

	entity.Sort = req.Sort

	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *exampleService) Update(req dto.ExampleUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.Example{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}

	entity.Name = req.Name
	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "example")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.Content = req.Content

	entity.Status = req.Status

	entity.Type = req.Type

	entity.IsVip = req.IsVip

	entity.Sort = req.Sort

	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

// 删除
func (s *exampleService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Example{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		if err != nil || rows == 0 {
			return 0, errors.New("删除失败")
		}
		return rows, nil
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.Example{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *exampleService) Status(req dto.ExampleStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.Example{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}

func (s *exampleService) IsVip(req dto.ExampleIsVipReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.Example{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.IsVip = req.IsVip
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}

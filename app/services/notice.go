/**
 * EasyBeeGo
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
	"strings"
	"time"
)

var Notice = new(noticeService)

type noticeService struct{}

func (s *noticeService) GetList(req dto.NoticePageReq) ([]vo.NoticeInfoVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Notice)).Filter("mark", 1)
	// 通知标题查询
	if req.Title != "" {
		query = query.Filter("title__contains", req.Title)
	}
	// 来源状态查询
	if req.Source > 0 {
		query = query.Filter("source", req.Source)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询总数
	count, _ := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.Notice, 0)
	query.All(&lists)

	// 数据处理
	var result []vo.NoticeInfoVo
	for _, v := range lists {
		item := vo.NoticeInfoVo{}
		item.Notice = v
		item.SourceName = constant.NOTICE_SOURCE_LIST[v.Source]
		result = append(result, item)
	}
	// 返回结果
	return result, count, nil
}

func (s *noticeService) Add(req dto.NoticeAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 富文本处理
	content := utils.SaveImageContent(req.Content, req.Title, "notice")
	// 实例化对象
	var entity models.Notice
	entity.Title = req.Title
	entity.Content = content
	entity.IsTop = req.IsTop
	entity.Source = req.Source
	entity.Status = req.Status
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入数据
	return entity.Insert()
}

func (s *noticeService) Update(req dto.NoticeUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.Notice{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}

	// 富文本处理
	content := utils.SaveImageContent(req.Content, req.Title, "notice")

	// 设置参数
	entity.Title = req.Title
	entity.Content = content
	entity.IsTop = req.IsTop
	entity.Source = req.Source
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

func (s *noticeService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Notice{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.Notice{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *noticeService) Status(req dto.NoticeStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.Notice{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新信息
	return entity.Update()
}

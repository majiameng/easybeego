/**
 * @author: Tinymeng <666@majiameng.com>
 */

package services

import (
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/conf"
	"easybeego/utils"
	"easybeego/utils/gconv"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"time"
)

// 中间件管理服务
var Level = new(levelService)

type levelService struct{}

func (s *levelService) GetList(req dto.LevelPageReq) ([]models.Level, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Level)).Filter("mark", 1)
	// 职级名称查询
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
	lists := make([]models.Level, 0)
	query.All(&lists)
	// 返回结果
	return lists, count, err
}

func (s *levelService) Add(req dto.LevelAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.Level
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

func (s *levelService) Update(req dto.LevelUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.Level{Id: req.Id}
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

func (s *levelService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 分裂记录ID
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Level{Id: gconv.Int(ids)}
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
			entity := &models.Level{Id: id}
			rows, err := entity.Delete()
			if rows == 0 || err != nil {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

// 设置状态
func (s *levelService) Status(req dto.LevelStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.Level{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}

func (s *levelService) ImportExcel(fileURL string, userId int) (int, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 获取本地文件绝对地址
	filePath := conf.CONFIG.Attachment.FilePath + fileURL
	// 读取Excel文件
	file, err := excelize.OpenFile(filePath)
	if err != nil {
		return 0, errors.New("excel文件读取失败")
	}
	// 读取第一张Sheet表
	rows, err := file.Rows("Sheet1")
	if err != nil {
		return 0, errors.New("excel文件读取失败")
	}
	// 计数器
	totalNum := 0
	// Excel文件头，此处必须与Excel模板头保持一致
	excelHeader := []string{"职级名称", "职级状态", "显示顺序"}
	// 循环遍历读取的数据源
	for rows.Next() {
		// Excel列对象
		item, err2 := rows.Columns()
		if err2 != nil {
			return 0, errors.New("excel文件解析异常")
		}
		// 读取的列数与Excel头列数不等则跳过读取下一条
		if len(item) != len(excelHeader) {
			continue
		}
		// 如果是标题栏则跳过
		if item[1] == "职级状态" {
			continue
		}
		// 职级名称
		name := item[0]
		// 职级状态
		status := 1
		if item[1] == "正常" {
			status = 1
		} else {
			status = 2
		}
		// 显示顺序
		sort, _ := strconv.Atoi(item[2])
		// 实例化职级导入对象
		level := models.Level{
			Name:       name,
			Status:     status,
			Sort:       sort,
			CreateUser: userId,
			CreateTime: time.Now(),
			UpdateUser: userId,
			UpdateTime: time.Now(),
			Mark:       1,
		}
		// 插入职级数据
		if _, err := level.Insert(); err != nil {
			return 0, err
		}
		// 计数器+1
		totalNum++
	}
	return totalNum, nil
}

func (s *levelService) GetExcelList(req dto.LevelPageReq) (string, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Level)).Filter("mark", 1)
	// 职级名称查询
	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询列表
	lists := make([]models.Level, 0)
	query.All(&lists)

	// 循环遍历处理数据源
	excel := excelize.NewFile()
	excel.SetSheetRow("Sheet1", "A1", &[]string{"ID", "职级名称", "职级状态", "排序", "创建时间"})
	for i, v := range lists {
		axis := fmt.Sprintf("A%d", i+2)
		excel.SetSheetRow("Sheet1", axis, &[]interface{}{
			v.Id,
			v.Name,
			v.Status,
			v.Sort,
			v.CreateTime,
		})
	}
	// 定义文件名称
	fileName := fmt.Sprintf("%s.xlsx", time.Now().Format("20060102150405"))
	// 设置Excel保存路径
	filePath := fmt.Sprintf("%s/temp/%s", conf.CONFIG.Attachment.FilePath, fileName)
	err2 := excel.SaveAs(filePath)
	// 获取文件URL地址
	fileURL := utils.GetImageUrl(strings.ReplaceAll(filePath, conf.CONFIG.Attachment.FilePath, ""))
	// 返回结果
	return fileURL, err2
}

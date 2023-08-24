package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Ad struct {
	Id          int       `orm:"column(id);auto"`
	Title       string    `orm:"column(title);size(100);null" description:"广告标题"`
	AdSortId    int       `orm:"column(ad_sort_id);null" description:"广告位ID"`
	Cover       string    `orm:"column(cover);size(100);null" description:"广告图片"`
	Type        int       `orm:"column(type);null" description:"广告格式：1图片 2文字 3视频 4推荐"`
	Description string    `orm:"column(description);size(100);null" description:"广告描述"`
	Content     string    `orm:"column(content);size(255);null" description:"广告内容"`
	Url         string    `orm:"column(url);size(255);null" description:"广告链接"`
	Width       int       `orm:"column(width);null" description:"广告宽度"`
	Height      int       `orm:"column(height);null" description:"广告高度"`
	StartTime   time.Time `orm:"column(start_time);type(datetime);null" description:"开始时间"`
	EndTime     time.Time `orm:"column(end_time);type(datetime);null" description:"结束时间"`
	Status      int       `orm:"column(status);null" description:"状态：1在用 2停用"`
	Sort        int       `orm:"column(sort);null" description:"排序"`
	Note        string    `orm:"column(note);size(255);null" description:"备注"`
	CreateUser  int       `orm:"column(create_user);null"`
	UpdateUser  int       `orm:"column(update_user);null"`
	Mark        int       `orm:"column(mark);null"`
	CreateTime  time.Time `orm:"column(create_time);type(datetime);null"`
	UpdateTime  time.Time `orm:"column(update_time);type(datetime);null"`
}

func (t *Ad) TableName() string {
	return "ad"
}

func init() {
	orm.RegisterModel(new(Ad))
}

// 根据条件查询单条数据
func (t *Ad) Get() error {
	err := orm.NewOrm().QueryTable(new(Ad)).Filter("id", t.Id).One(t)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return errors.New("查询到了多条记录")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return errors.New("未查询到记录")
	}
	return nil
}

// 插入数据
func (t *Ad) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Ad) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Ad) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

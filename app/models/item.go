package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Item struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(255);null" description:"站点类型:1普通站点 2其他"`
	Type       int       `orm:"column(type);null" description:"站点类型:1普通站点 2其他"`
	Url        string    `orm:"column(url);size(255);null" description:"站点地址"`
	Image      string    `orm:"column(image);size(255);null" description:"站点图片"`
	Status     int       `orm:"column(status);null" description:"状态：1在用 2停用"`
	Sort       int       `orm:"column(sort);null" description:"排序"`
	Note       string    `orm:"column(note);size(255);null" description:"备注"`
	Mark       int       `orm:"column(mark);null"`
	CreateUser int       `orm:"column(create_user);null" description:"创建人"`
	UpdateUser int       `orm:"column(update_user);null" description:"修改人"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"修改时间"`
}

func (t *Item) TableName() string {
	return "item"
}

func init() {
	orm.RegisterModel(new(Item))
}

// 根据条件查询单条数据
func (t *Item) Get() error {
	err := orm.NewOrm().QueryTable(new(Item)).Filter("id", t.Id).One(t)
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
func (t *Item) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Item) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Item) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Example struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(255);null" description:"测试名称"`
	Avatar     string    `orm:"column(avatar);size(255);null" description:"头像"`
	Content    string    `orm:"column(content);size(255);null" description:"内容"`
	Status     int       `orm:"column(status);null" description:"状态：1正常 2停用"`
	Type       int       `orm:"column(type);null" description:"类型：1京东 2淘宝 3拼多多 4唯品会"`
	IsVip      int       `orm:"column(is_vip);null" description:"是否VIP：1是 2否"`
	Sort       int       `orm:"column(sort);null" description:"排序号"`
	Mark       int       `orm:"column(mark);null"`
	CreateUser int       `orm:"column(create_user);null" description:"创建人"`
	UpdateUser int       `orm:"column(update_user);null" description:"修改人"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"修改时间"`
}

func (t *Example) TableName() string {
	return "example"
}

func init() {
	orm.RegisterModel(new(Example))
}

// 根据条件查询单条数据
func (t *Example) Get() error {
	err := orm.NewOrm().QueryTable(new(Example)).Filter("id", t.Id).One(t)
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
func (t *Example) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Example) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Example) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Role struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(100);null" description:"广告标题"`
	Code       string    `orm:"column(code);size(10);null" description:"广告位ID"`
	Status     int       `orm:"column(status);null" description:"状态：1在用 2停用"`
	Sort       int       `orm:"column(sort);null" description:"排序"`
	Note       string    `orm:"column(note);size(255);null" description:"备注"`
	Mark       int       `orm:"column(mark);null"`
	CreateUser int       `orm:"column(create_user);null"`
	UpdateUser int       `orm:"column(update_user);null"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null"`
}

func (t *Role) TableName() string {
	return "role"
}

func init() {
	orm.RegisterModel(new(Role))
}

// 根据条件查询单条数据
func (t *Role) Get() error {
	err := orm.NewOrm().QueryTable(new(Role)).Filter("id", t.Id).One(t)
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
func (t *Role) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Role) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Role) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

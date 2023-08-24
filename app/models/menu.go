package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Menu struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(100);null" description:"菜单标题"`
	Icon       string    `orm:"column(icon);size(10);null" description:"图标"`
	Url        string    `orm:"column(url);size(100);null" description:"URL地址"`
	Pid        int       `orm:"column(pid);null" description:"上级ID"`
	Type       int       `orm:"column(type);null" description:"类型：1模块 2导航 3菜单 4节点"`
	Permission string    `orm:"column(permission);size(255);null" description:"权限标识"`
	Status     int       `orm:"column(status);null" description:"状态：1正常 2禁用"`
	Target     int       `orm:"column(target);null" description:"打开方式：1内部打开 2外部打开"`
	Sort       int       `orm:"column(sort);null" description:"排序"`
	Note       string    `orm:"column(note);size(255);null" description:"备注"`
	Func       string    `orm:"column(func);size(255);null" description:"权限节点"`
	Mark       int       `orm:"column(mark);null"`
	CreateUser int       `orm:"column(create_user);null"`
	UpdateUser int       `orm:"column(update_user);null"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null"`
}

func (t *Menu) TableName() string {
	return "menu"
}

func init() {
	orm.RegisterModel(new(Menu))
}

// 根据条件查询单条数据
func (t *Menu) Get() error {
	err := orm.NewOrm().QueryTable(new(Menu)).Filter("id", t.Id).One(t)
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
func (t *Menu) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Menu) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Menu) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

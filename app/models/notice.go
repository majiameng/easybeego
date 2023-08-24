package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Notice struct {
	Id         int       `orm:"column(id);auto"`
	Title      string    `orm:"column(title);size(100);null" description:"通知标题"`
	Content    string    `orm:"column(content);size(10);null" description:"通知内容"`
	Source     int       `orm:"column(source);null" description:"来源：1内部通知 2外部新闻"`
	IsTop      int       `orm:"column(is_top);null" description:"是否置顶：1是 2否"`
	Status     int       `orm:"column(status);null" description:"状态：1已发布 2待发布"`
	Sort       int       `orm:"column(sort);null" description:"排序"`
	Mark       int       `orm:"column(mark);null" description:"备注"`
	CreateUser int       `orm:"column(create_user);null"`
	UpdateUser int       `orm:"column(update_user);null"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null"`
}

func (t *Notice) TableName() string {
	return "notice"
}

func init() {
	orm.RegisterModel(new(Notice))
}

// 根据条件查询单条数据
func (t *Notice) Get() error {
	err := orm.NewOrm().QueryTable(new(Notice)).Filter("id", t.Id).One(t)
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
func (t *Notice) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *Notice) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *Notice) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

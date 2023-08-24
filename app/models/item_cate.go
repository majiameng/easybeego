package models

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type ItemCate struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(255);null" description:"栏目名称"`
	Pid        int       `orm:"column(pid);null" description:"父级ID"`
	ItemId     int       `orm:"column(item_id);null" description:"栏目ID"`
	Pinyin     string    `orm:"column(pinyin);size(255);null" description:"拼音(全)"`
	Code       string    `orm:"column(code);size(255);null" description:"拼音(简)"`
	IsCover    int       `orm:"column(is_cover);null" description:"是否有封面：1是 2否"`
	Cover      string    `orm:"column(cover);size(255);null" description:"封面"`
	Status     int       `orm:"column(status);null" description:"状态：1启用 2停用"`
	Sort       int       `orm:"column(sort);null" description:"排序"`
	Note       string    `orm:"column(note);size(255);null" description:"备注"`
	Mark       int       `orm:"column(mark);null"`
	CreateUser int       `orm:"column(create_user);null" description:"创建人"`
	UpdateUser int       `orm:"column(update_user);null" description:"修改人"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"修改时间"`
}

func (t *ItemCate) TableName() string {
	return "item_cate"
}

func init() {
	orm.RegisterModel(new(ItemCate))
}

// 根据条件查询单条数据
func (t *ItemCate) Get() error {
	err := orm.NewOrm().QueryTable(new(ItemCate)).Filter("id", t.Id).One(t)
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
func (t *ItemCate) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *ItemCate) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *ItemCate) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type City struct {
	Id         int       `orm:"column(id);auto"`
	Pid        int       `orm:"column(pid);null" description:"父级编号"`
	Level      int       `orm:"column(level);null" description:"城市级别：1省 2市 3区"`
	Name       string    `orm:"column(name);size(255);null" description:"城市名称"`
	Citycode   string    `orm:"column(citycode);size(255);null" description:"城市编号（区号）"`
	PAdcode    string    `orm:"column(p_adcode);size(255);null" description:"父级地理编号"`
	Adcode     string    `orm:"column(adcode);size(255);null" description:"地理编号"`
	Lng        string    `orm:"column(lng);size(255);null" description:"城市坐标中心点经度（* 1e6）：如果是中国，此值是 1e7"`
	Lat        string    `orm:"column(lat);size(255);null" description:"城市坐标中心点纬度（* 1e6）"`
	Sort       int       `orm:"column(sort);null" description:"排序号"`
	Mark       int       `orm:"column(mark);null" description:""`
	CreateUser int       `orm:"column(create_user);null" description:"创建人"`
	UpdateUser int       `orm:"column(update_user);null" description:"修改人"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime);null" description:"修改时间"`
}

func (t *City) TableName() string {
	return "city"
}

func init() {
	orm.RegisterModel(new(City))
}

// 根据条件查询单条数据
func (t *City) Get() error {
	err := orm.NewOrm().QueryTable(new(City)).Filter("id", t.Id).One(t)
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
func (t *City) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *City) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *City) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

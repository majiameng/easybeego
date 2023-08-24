package models

import (
	"errors"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type User struct {
	Id           int       `orm:"column(id);auto"`
	Realname     string    `orm:"column(realname);size(255);null" description:"真实姓名"`
	Nickname     string    `orm:"column(nickname);size(255);null" description:"用户昵称"`
	Gender       int       `orm:"column(gender);null" description:"性别（1男 2女 3未知）"`
	Avatar       string    `orm:"column(avatar);size(255);null" description:"用户头像"`
	Mobile       string    `orm:"column(mobile);size(255);null" description:"手机号"`
	Email        string    `orm:"column(email);size(255);null" description:"邮箱"`
	Birthday     time.Time `orm:"column(birthday);type(datetime);null" description:"出生日期"`
	DeptId       int       `orm:"column(dept_id);null"`
	LevelId      int       `orm:"column(level_id);null"`
	PositionId   int       `orm:"column(position_id);null"`
	ProvinceCode string    `orm:"column(province_code);size(255);null" description:"省份编号"`
	CityCode     string    `orm:"column(city_code);size(255);null" description:"市区编号"`
	DistrictCode string    `orm:"column(district_code);size(255);null" description:"区县编号"`
	Address      string    `orm:"column(address);size(255);null" description:"详细地址"`
	Username     string    `orm:"column(username);size(255);null" description:"用户名"`
	Password     string    `orm:"column(password);size(255);null" description:"登录密码"`
	Intro        string    `orm:"column(intro);size(255);null" description:"个人简介"`
	Status       int       `orm:"column(status);null" description:"是否启用 1、启用  2、停用"`
	Note         string    `orm:"column(note);size(255);null"`
	Sort         int       `orm:"column(sort);null"`
	LoginTime    time.Time `orm:"column(login_time);type(datetime);null"`
	LoginIp      string    `orm:"column(login_ip);size(255);null"`
	RoleIds      string    `orm:"column(role_ids);size(255);null" description:"用户角色"`
	Mark         int       `orm:"column(mark);null"`
	CreateUser   int       `orm:"column(create_user);null" description:"创建人"`
	UpdateUser   int       `orm:"column(update_user);null" description:"修改人"`
	CreateTime   time.Time `orm:"column(create_time);type(datetime);null" description:"创建时间"`
	UpdateTime   time.Time `orm:"column(update_time);type(datetime);null" description:"修改时间"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// 根据条件查询单条数据
func (t *User) Get() error {
	err := orm.NewOrm().QueryTable(new(User)).Filter("id", t.Id).One(t)
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
func (t *User) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(t)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 更新数据
func (t *User) Update() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Update(t)
	if rows == 0 || err != nil {
		return 0, errors.New("更新失败")
	}
	return rows, nil
}

// 删除记录
func (t *User) Delete() (int64, error) {
	o := orm.NewOrm()
	rows, err := o.Delete(t)
	if rows == 0 || err != nil {
		return 0, errors.New("删除失败")
	}
	return rows, nil
}

/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package services

import (
	"easybeego/app/constant"
	"easybeego/app/dto"
	"easybeego/app/models"
	"easybeego/app/vo"
	"easybeego/utils"
	"easybeego/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"time"
)

var User = new(userService)

type userService struct{}

func (s *userService) GetList(req dto.UserPageReq) ([]vo.UserInfoVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.User)).Filter("mark", 1)
	// 用户名
	if req.Realname != "" {
		query = query.Filter("realname__contains", req.Realname)
	}
	// 性别
	if req.Gender > 0 {
		query = query.Filter("gender", req.Gender)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询总数
	count, err := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.User, 0)
	query.All(&lists)

	// 获取职级列表
	levelList := make([]models.Level, 0)
	orm.NewOrm().QueryTable(new(models.Level)).Filter("mark", 1).OrderBy("sort").All(&levelList)
	var levelMap = make(map[int]string)
	for _, v := range levelList {
		levelMap[v.Id] = v.Name
	}
	// 获取岗位列表
	positionList := make([]models.Position, 0)
	orm.NewOrm().QueryTable(new(models.Position)).Filter("mark", 1).OrderBy("sort").All(&positionList)
	var positionMap = make(map[int]string)
	for _, v := range positionList {
		positionMap[v.Id] = v.Name
	}
	// 获取部门列表
	deptList := make([]models.Dept, 0)
	orm.NewOrm().QueryTable(new(models.Dept)).Filter("mark", 1).OrderBy("sort").All(&deptList)
	var deptMap = make(map[int]string)
	for _, v := range deptList {
		deptMap[v.Id] = v.Name
	}
	// 数据处理
	var result []vo.UserInfoVo
	for _, v := range lists {
		item := vo.UserInfoVo{}
		item.User = v
		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}
		// 性别
		if v.Gender > 0 {
			item.GenderName = constant.GENDER_LIST[v.Gender]
		}
		// 职级
		if v.LevelId > 0 {
			item.LevelName = levelMap[v.LevelId]
		}
		// 岗位
		if v.PositionId > 0 {
			item.PositionName = positionMap[v.PositionId]
		}
		// 部门
		if v.DeptId > 0 {
			item.DeptName = deptMap[v.DeptId]
		}
		// 角色列表
		roleList := UserRole.GetUserRoleList(v.Id)
		if len(roleList) > 0 {
			item.RoleList = roleList
		} else {
			item.RoleList = make([]models.Role, 0)
		}
		// 省市区
		cityList := make([]string, 0)
		// 省份编号
		cityList = append(cityList, item.ProvinceCode)
		// 城市编号
		cityList = append(cityList, item.CityCode)
		// 县区编号
		cityList = append(cityList, item.DistrictCode)
		item.City = cityList
		// 加入数组
		result = append(result, item)
	}
	// 返回结果
	return result, count, err
}

func (s *userService) Add(req dto.UserAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	var entity models.User
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	entity.Avatar = req.Avatar
	entity.Mobile = req.Mobile
	entity.Email = req.Email
	// 日期处理
	tm2, _ := time.Parse("2006-01-02", req.Birthday)
	entity.Birthday = tm2
	entity.DeptId = req.DeptId
	entity.LevelId = req.LevelId
	entity.PositionId = req.PositionId
	entity.ProvinceCode = req.ProvinceCode
	entity.CityCode = req.CityCode
	entity.DistrictCode = req.DistrictCode
	entity.Address = req.Address
	entity.Username = req.Username
	entity.Intro = req.Intro
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		entity.Password = password
	}

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1

	// 插入记录
	rows, err := entity.Insert()
	if err != nil || rows == 0 {
		return 0, err
	}

	// 删除用户角色关系
	orm.NewOrm().QueryTable(new(models.UserRole)).Filter("user_id", entity.Id).Delete()
	// 创建人员角色关系
	roleIdArr := strings.Split(req.RoleIds, ",")
	for _, v := range roleIdArr {
		if v == "" {
			continue
		}
		var userRole models.UserRole
		userRole.UserId = entity.Id
		userRole.RoleId = gconv.Int(v)
		userRole.Insert()
	}
	return rows, nil
}

func (s *userService) Update(req dto.UserUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	entity := &models.User{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	entity.Avatar = req.Avatar
	entity.Mobile = req.Mobile
	entity.Email = req.Email
	// 日期处理
	tm2, _ := time.Parse("2006-01-02", req.Birthday)
	entity.Birthday = tm2
	entity.DeptId = req.DeptId
	entity.LevelId = req.LevelId
	entity.PositionId = req.PositionId
	entity.ProvinceCode = req.ProvinceCode
	entity.CityCode = req.CityCode
	entity.DistrictCode = req.DistrictCode
	entity.Address = req.Address
	entity.Username = req.Username
	entity.Intro = req.Intro
	entity.Status = req.Status
	entity.Note = req.Note
	entity.Sort = req.Sort

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		entity.Password = password
	}

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "user")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	rows, err := entity.Update()
	if err != nil || rows == 0 {
		return 0, err
	}

	// 删除用户角色关系
	orm.NewOrm().QueryTable(new(models.UserRole)).Filter("user_id", entity.Id).Delete()
	// 创建人员角色关系
	roleIdArr := strings.Split(req.RoleIds, ",")
	for _, v := range roleIdArr {
		if v == "" {
			continue
		}
		var userRole models.UserRole
		userRole.UserId = entity.Id
		userRole.RoleId = gconv.Int(v)
		userRole.Insert()
	}
	return rows, nil
}

func (s *userService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.User{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.User{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *userService) Status(req dto.UserStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.User{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}

func (s *userService) UpdateUserInfo(req dto.UserInfoReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 获取用户信息
	user := &models.User{Id: userId}
	err := user.Get()
	if err != nil {
		return 0, err
	}
	//// 头像处理
	//if req.Avatar != "" {
	//	avatar, err := utils.SaveImage(req.Avatar, "user")
	//	if err != nil {
	//		return 0, err
	//	}
	//	user.Avatar = avatar
	//}
	user.Realname = req.Realname
	user.Nickname = req.Nickname
	user.Gender = req.Gender
	user.Mobile = req.Mobile
	user.Email = req.Email
	user.Address = req.Address
	user.Intro = req.Intro
	user.UpdateUser = userId
	user.UpdateTime = time.Now()
	rows, err := user.Update()
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func (s *userService) UpdatePwd(req dto.UpdatePwd, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询信息
	user := &models.User{Id: userId}
	err := user.Get()
	if err != nil {
		return 0, err
	}
	if user == nil {
		return 0, errors.New("记录不存在")
	}
	// 比对旧密码
	oldPwd, err := utils.Md5(req.OldPassword + user.Username)
	if err != nil {
		return 0, err
	}
	if oldPwd != user.Password {
		return 0, errors.New("旧密码不正确")
	}

	// 设置新密码
	if req.NewPassword != req.RePassword {
		return 0, errors.New("两次输入的新密码不一致")
	}
	newPwd, err := utils.Md5(req.NewPassword + user.Username)
	if err != nil {
		return 0, err
	}

	// 指定字段更新
	o := orm.NewOrm()
	entity := models.User{Id: userId}
	if o.Read(&entity) == nil {
		entity.Password = newPwd
		entity.UpdateUser = userId
		entity.UpdateTime = time.Now()
		return o.Update(&entity, "Password", "UpdateUser", "UpdateTime")
	}
	return 0, nil
}

func (s *userService) ResetPwd(id int, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	info := &models.User{Id: id}
	err := info.Get()
	if err != nil || info == nil {
		return 0, err
	}
	// 设置初始密码
	password, err := utils.Md5("123456" + info.Username)
	if err != nil {
		return 0, err
	}

	// 初始化密码
	o := orm.NewOrm()
	user := models.User{Id: id}
	if o.Read(&user) == nil {
		user.Password = password
		user.UpdateUser = userId
		user.UpdateTime = time.Now()
		return o.Update(&user, "Password", "UpdateUser", "UpdateTime")
	}
	return 0, nil
}

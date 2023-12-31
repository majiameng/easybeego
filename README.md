## 📚 EasyBeeGo 项目介绍
一款 Go 语言基于Beego、Layui、MySQL等框架精心打造的一款模块化、高性能、企业级的敏捷开发框架，本着简化开发、提升开发效率的初衷触发，框架自研了一套个性化的组件，实现了可插拔的组件式开发方式：单图上传、多图上传、下拉选择、开关按钮、单选按钮、多选按钮、图片裁剪等等一系列个性化、轻量级的组件，是一款真正意义上实现组件化开发的敏捷开发框架。

### 生成代码
bee generate appcode -driver=mysql -conn="easybeego:MbKwf2fR7Xii74Zz@tcp(127.0.0.1:3306)/easybeego" -level=1

### beego框架：bee工具进行beego项目的创建、热编译、开发、测试、和部署
https://sns.bjwmsc.com/archives/3813

## 🍻 项目特点

+ 模块化、松耦合
+ 模块丰富、开箱即用
+ 简洁易用、快速接入
+ 文档详尽、易于维护
+ 自顶向下、体系化设计
+ 统一框架、统一组件、降低选择成本
+ 开发规范、设计模式、代码分层模型
+ 强大便捷的开发工具链
+ 完善的本地中文化支持
+ 设计为团队及企业使用

## 🍪 内置模块
+ 用户管理：用于维护管理系统的用户，常规信息的维护与账号设置。
+ 角色管理：角色菜单管理与权限分配、设置角色所拥有的菜单权限。
+ 菜单管理：配置系统菜单，操作权限，按钮权限标识等。
+ 职级管理：主要管理用户的职级。
+ 岗位管理：主要管理用户担任职务。
+ 部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。
+ 字典管理：对系统中常用的较为固定的数据进行统一维护。
+ 配置管理：对系统的常规配置信息进行维护，网站配置管理功能进行统一维护。
+ 通知公告：系统通知公告信息发布维护。
+ 操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
+ 登录日志：系统登录日志记录查询包含登录异常。
+ 代码生成：一键生成模块CRUD的功能，包括后端Go和前端HTML、JS等相关代码。
+ 案例演示：常规代码生成器一键生成后的演示案例。

## 🍪 项目结构

```
EasyBeeGo
|-- conf
|   `-- app.conf
|-- controllers
|   `-- default.go
|-- main.go
|-- models
|-- routers
|   `-- router.go
|-- static
|   |-- css
|   |-- img
|   `-- js
|-- tests
|   `-- default_test.go
`-- views
    `-- index.tpl
```

## 📚 核心组件

+ 单图上传组件
```
{{upload_image "avatar|头像|90x90|建议上传尺寸450x450|450x450" .info.Avatar "" 0}}
```
+ 多图上传组件
```
{{album "avatar|图集|90x90|20|建议上传尺寸450x450" .info.Avatar "" 0}}
```
+ 下拉选择组件
```
{{select "gender|1|性别|name|id" "1=男,2=女,3=保密" .info.Gender}}
```
+ 单选按钮组件
```
{{radio "gender|name|id" "1=男,2=女,3=保密" .info.Gender}}
```
+ 复选框组件
```
{{checkbox "role_ids|name|id" .roleList .info.RoleIds}}
```
+ 城市选择组件
```
{{city .info.DistrictCode 3 1}}
```
+ 开关组件
```
{{switch "status" "在用|禁用" .info.Status}}
```
+ 日期组件
```
{{date "birthday|1|出生日期|date" .info.Birthday}}
```
+ 图标组件
```
{{icon "icon" .info.Icon}}
```
+ 穿梭组件
```
{{transfer "func|0|全部节点,已赋予节点|name|id|220x350" "1=列表,5=添加,10=修改,15=删除,20=详情,25=状态,30=批量删除,35=添加子级,40=全部展开,45=全部折叠" .funcList}}
```

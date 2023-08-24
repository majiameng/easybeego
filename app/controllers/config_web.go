/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/models"
	"easybeego/conf"
	"easybeego/utils"
	"easybeego/utils/common"
	"easybeego/utils/gconv"
	"easybeego/utils/gstr"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"regexp"
	"strings"
	"time"
)

var ConfigWeb = new(ConfigWebController)

type ConfigWebController struct {
	BaseController
}

func (ctl *ConfigWebController) Index() {

	if ctl.Ctx.Input.IsPost() {
		// 返回结果
		if utils.AppDebug() {
			ctl.JSON(common.JsonResult{
				Code: -1,
				Msg:  "演示环境，暂无权限操作",
			})
		}
		// key：string类型，value：interface{}  类型能存任何数据类型
		var jsonObj map[string]interface{}
		// //在RequestBody中读取Json
		data := ctl.Ctx.Input.RequestBody
		json.Unmarshal(data, &jsonObj)

		// 遍历处理数据源
		for key, val := range jsonObj {
			// 参数处理
			if key == "checkbox" {
				// 复选框

				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
			} else if strings.Contains(key, "upimage") {
				// 单图上传

				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
				if strings.Contains(gconv.String(val), "temp") {
					image, _ := utils.SaveImage(gconv.String(val), "config")
					// 赋值给参数
					val = image
				} else {
					// 赋值给参数
					val = gstr.Replace(gconv.String(val), conf.CONFIG.EGAdmin.Image, "")
				}
			} else if strings.Contains(key, "upimgs") {
				// 多图上传
				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]
				// 图片地址处理
				urlArr := gstr.Split(gconv.String(val), ",")
				list := make([]string, 0)
				for _, v := range urlArr {
					if strings.Contains(gconv.String(v), "temp") {
						image, _ := utils.SaveImage(v, "config")
						list = append(list, image)
					} else {
						image := gstr.Replace(v, conf.CONFIG.EGAdmin.Image, "")
						list = append(list, image)
					}
				}
				// 数组转字符串，逗号分隔
				val = strings.Join(list, ",")
			} else if strings.Contains(key, "ueditor") {
				item := gstr.Split(key, "__")
				// KEY值
				key = item[0]

				// 富文本处理(待完善)
				// TODO...
			}

			var entity models.ConfigData
			err := orm.NewOrm().QueryTable(new(models.ConfigData)).Filter("code", key).One(&entity)
			if err != nil {
				continue
			}
			// 更新记录
			entity.Value = gconv.String(val)
			entity.UpdateUser = utils.Uid(ctl.Ctx)
			entity.UpdateTime = time.Now()
			entity.Update()
		}

		// 返回结果
		ctl.JSON(common.JsonResult{
			Code: 0,
			Msg:  "保存成功",
		})
	}
	// 配置ID
	configId := ctl.GetString("configId", "")
	if configId == "" {
		configId = "1"
	}

	// 获取配置列表
	configData := make([]models.Config, 0)
	orm.NewOrm().QueryTable(new(models.Config)).Filter("mark", 1).All(&configData)
	configList := make(map[string]string)
	for _, v := range configData {
		configList[gconv.String(v.Id)] = v.Name
	}

	// 获取配置项列表
	itemData := make([]models.ConfigData, 0)
	orm.NewOrm().QueryTable(new(models.ConfigData)).
		Filter("config_id", configId).
		Filter("status", 1).
		Filter("mark", 1).
		OrderBy("sort").
		All(&itemData)
	itemList := make([]map[string]interface{}, 0)
	for _, v := range itemData {
		item := make(map[string]interface{})
		item["id"] = v.Id
		item["title"] = v.Title
		item["code"] = v.Code
		item["value"] = v.Value
		item["type"] = v.Type

		if v.Type == "checkbox" {
			// 复选框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, val := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(val, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList
		} else if v.Type == "radio" {
			// 单选框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, v := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList

		} else if v.Type == "select" {
			// 下拉选择框
			re := regexp.MustCompile(`\r?\n`)
			list := gstr.Split(re.ReplaceAllString(v.Options, "|"), "|")
			optionsList := make(map[string]string)
			for _, v := range list {
				re2 := regexp.MustCompile(`:|：|\s+`)
				item := gstr.Split(re2.ReplaceAllString(v, "|"), "|")
				optionsList[item[0]] = item[1]
			}
			item["optionsList"] = optionsList
		} else if v.Type == "image" {
			// 单图片
			item["value"] = utils.GetImageUrl(v.Value)
		} else if v.Type == "images" {
			// 多图片
			list := gstr.Split(v.Value, ",")
			itemList := make([]string, 0)
			for _, v := range list {
				// 图片地址
				item := utils.GetImageUrl(v)
				itemList = append(itemList, item)
			}
			item["value"] = itemList
		}
		itemList = append(itemList, item)
	}

	// 渲染模板
	ctl.Data["configId"] = configId
	ctl.Data["configList"] = configList
	ctl.Data["itemList"] = itemList
	ctl.Layout = "public/layout.html"
	ctl.TplName = "config_web/index.html"
}

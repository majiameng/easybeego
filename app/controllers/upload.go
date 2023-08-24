/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

import (
	"easybeego/app/services"
	"easybeego/utils"
	"easybeego/utils/common"
)

// 控制器管理对象
var Upload = new(UploadController)

type UploadController struct {
	BaseController
}

func (ctl *UploadController) UploadImage() {
	// 调用上传方法
	result, err := services.Upload.UploadImage(ctl.Ctx)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 拼接图片地址
	result.FileUrl = utils.GetImageUrl(result.FileUrl)
	// 返回结果
	ctl.JSON(common.JsonResult{
		Code: 0,
		Msg:  "上传成功",
		Data: result,
	})
}

func (ctl *UploadController) UploadEditImage() {
	// 调用上传方法
	result, err := services.Upload.UploadImage(ctl.Ctx)
	if err != nil {
		ctl.JSON(common.JsonResult{
			Code: -1,
			Msg:  err.Error(),
		})
	}
	// 拼接图片地址
	fileUrl := utils.GetImageUrl(result.FileUrl)
	// 返回结果
	ctl.JSON(common.JsonEditResult{
		Error: 0,
		Url:   fileUrl,
	})
}

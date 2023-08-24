/**
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

var Generate = new(GenerateController)

type GenerateController struct {
	BaseController
}

//
//func (ctl *GenerateController) Index() {
//	// 渲染模板
//	ctl.Layout = "public/layout.html"
//	ctl.TplName = "generate/index.html"
//}
//
//func (ctl *GenerateController) List() {
//	// 参数验证
//	var req dto.GeneratePageReq
//	if err := ctl.ParseForm(&req); err != nil {
//		ctl.JSON(common.JsonResult{
//			Code: -1,
//			Msg:  err.Error(),
//		})
//		return
//	}
//
//	// 调用查询列表方法
//	list, err := services.Generate.GetList(req)
//	if err != nil {
//		ctl.JSON(common.JsonResult{
//			Code: -1,
//			Msg:  err.Error(),
//		})
//		return
//	}
//
//	// 返回结果
//	ctl.JSON(common.JsonResult{
//		Code:  0,
//		Msg:   "查询成功",
//		Data:  list,
//		Count: gconv.Int64(len(list)),
//	})
//}
//
//func (ctl *GenerateController) Generate() {
//	// 生成对象
//	var req dto.GenerateFileReq
//	if err := ctl.ParseForm(&req); err != nil {
//		ctl.JSON(common.JsonResult{
//			Code: -1,
//			Msg:  err.Error(),
//		})
//	}
//	// 参数校验
//	v := validate.Struct(req)
//	if !v.Validate() {
//		ctl.JSON(common.JsonResult{
//			Code: -1,
//			Msg:  v.Errors.One(),
//		})
//	}
//	// 调用生成方法
//	err := services.Generate.Generate(req, ctl.Ctx)
//	if err != nil {
//		ctl.JSON(common.JsonResult{
//			Code: -1,
//			Msg:  err.Error(),
//		})
//		return
//	}
//	// 返回结果
//	ctl.JSON(common.JsonResult{
//		Code: 0,
//		Msg:  "生成成功",
//	})
//}
//
//func (ctl *GenerateController) BatchGenerate() {
//	// 生成对象
//	var req dto.BatchGenerateFileReq
//	if err := ctl.ParseForm(&req); err != nil {
//		ctl.JSON(common.JsonResult{
//			Code: -1,
//			Msg:  err.Error(),
//		})
//	}
//	// 参数校验
//	v := validate.Struct(req)
//	if !v.Validate() {
//		ctl.JSON(common.JsonResult{
//			Code: -1,
//			Msg:  v.Errors.One(),
//		})
//	}
//	// 参数分析
//	tableList := strings.Split(req.Tables, ",")
//	count := 0
//	for _, item := range tableList {
//		itemList := strings.Split(item, "|")
//		// 组装参数对象
//		var param dto.GenerateFileReq
//		param.Name = itemList[0]
//		param.Comment = itemList[1]
//		// 调用生成方法
//		err := services.Generate.Generate(param, ctl.Ctx)
//		if err != nil {
//			continue
//		}
//		count++
//	}
//	// 返回结果
//	ctl.JSON(common.JsonResult{
//		Code: 0,
//		Msg:  "本次共生成【" + strconv.Itoa(count) + "】个模块文件",
//	})
//}

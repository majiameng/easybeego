/**
 * @author: Tinymeng <666@majiameng.com>
 */

package controllers

var Analysis = new(AnalysisController)

type AnalysisController struct {
	BaseController
}

func (ctl *AnalysisController) Index() {
	ctl.Layout = "public/layout.html"
	ctl.TplName = "analysis/index.html"
}

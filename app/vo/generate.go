/**
 * @author: Tinymeng <666@majiameng.com>
 */

package vo

// 数据库信息
type GenerateInfo struct {
	Name          string `json:"Name"`          // 表名
	Engine        string `json:"Engine"`        // 引擎
	Version       string `json:"Version"`       // 版本
	Collation     string `json:"Collation"`     // 编码
	Rows          string `json:"Rows"`          // 记录数
	DataLength    string `json:"DataLength"`    // 大小
	AutoIncrement string `json:"AutoIncrement"` // 自增索引
	Comment       string `json:"Comment"`       // 表备注
	CreateTime    string `json:"CreateTime"`    // 添加时间
	UpdateTime    string `json:"UpdateTime"`    // 更新时间
}

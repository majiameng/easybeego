/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package config

// 上传附件配置
type Attachment struct {
	FilePath string `mapstructure:"file_path" json:"file_path" yaml:"file_path"` //上传目录配置（相对于根目录）
	FileSize string `mapstructure:"file_size" json:"file_size" yaml:"file_size"` //默认不超过50mb
	FileExt  string `mapstructure:"file_ext" json:"file_ext" yaml:"file_ext"`    //url（相对于web目录）
}

/**
 * EasyBeeGo
 * @author: Tinymeng <666@majiameng.com>
 */

package config

// 系统配置
type EGAdmin struct {
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	Debug   bool   `mapstructure:"debug" json:"debug" yaml:"debug"`
	Image   string `mapstructure:"image" json:"image" yaml:"image"`
	Uploads string `mapstructure:"uploads" json:"uploads" yaml:"uploads"`
}

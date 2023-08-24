/**
 * @author: Tinymeng <666@majiameng.com>
 */

package conf

import "easybeego/conf/config"

// 全局配置结构体
type Config struct {
	Mysql      config.MySQL      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Attachment config.Attachment `mapstructure:"attachment" json:"attachment" yaml:"attachment"`
	EGAdmin    config.EGAdmin    `mapstructure:"easybeego" json:"easybeego" yaml:"easybeego"`
}

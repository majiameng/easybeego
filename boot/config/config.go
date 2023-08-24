/**
 * @author: Tinymeng <666@majiameng.com>
 */

package config

import (
	"easybeego/conf"
	"easybeego/constant"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

// 加载配置文件
func init() {
	var config string
	if configEnv := os.Getenv(constant.ConfigEnv); configEnv == "" {
		config = constant.ConfigFile
		fmt.Printf("您正在使用config的默认值,config的路径为%v\n", constant.ConfigFile)
	} else {
		config = configEnv
		fmt.Printf("您正在使用CONFIG环境变量,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&conf.CONFIG); err != nil {
			panic(err)
		}
	})

	if err := v.Unmarshal(&conf.CONFIG); err != nil {
		panic(err)
	}
}

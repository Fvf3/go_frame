package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigFile("config") //配置文件名称
	viper.SetConfigType("yaml")   //配置文件类型
	viper.AddConfigPath(".")      //配置查找配置文件的路径
	err = viper.ReadInConfig()    // 读取配置文件
	if err != nil {
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err) //没读到
		return
	}
	viper.WatchConfig()                            //启用支持热加载
	viper.OnConfigChange(func(in fsnotify.Event) { //回调，当配置文件发生修改时触发
		fmt.Println("Config file changed:", in.Name)
	})
	return
}

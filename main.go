package main

import (
	"fmt"
	"go_frame/settings"
)

// go web通用开发模板
func main() {
	//1.加载配置
	if err := settings.Init(); err != nil {
		fmt.Printf("settings init error:%s", err.Error())
		return
	}
	//2.初始化日志
	if err := logger.Init(); err != nil {
		fmt.Printf("logger init error:%s", err.Error())
	}
	//3.初始化持久化数据库连接
	if err := mysql.Init(); err != nil {
		fmt.Printf("mysql init error:%s", err.Error())
	}
	//4.初始化缓存数据库连接
	if err := redis.Init(); err != nil {
		fmt.Printf("redis init error:%s", err.Error())
	}
	//5.注册路由
	//6.启动服务（优雅关机)
}

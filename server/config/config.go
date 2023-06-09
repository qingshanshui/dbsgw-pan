package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	var mode string
	flag.StringVar(&mode, "mode", "dev", "开发dev，生产prod") // 开发环境
	flag.Parse()
	workDir, _ := os.Getwd()
	viper.SetConfigName("config." + mode)    // 配置文件名称
	viper.SetConfigType("yml")               // 配置文件类型
	viper.AddConfigPath(workDir + "/config") // 配置文件路径
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 动态加载配置文件
	viper.WatchConfig()
	// 动态加载配置文件事件
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置发生变化的文件：", e.Name)
	})
}

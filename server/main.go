package main

import (
	_ "fiber-layout/config"
	"fiber-layout/initalize"
	_ "fiber-layout/initalize"
	"fiber-layout/models"
	"fiber-layout/routers"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
)

func main() {
	app := fiber.New()
	// 静态目录
	app.Static("/", "./static")
	// HTTP 请求/响应日志
	app.Use(logger.New())
	// 使用各种选项启用跨源资源共享(CORS)
	app.Use(cors.New())
	// Recover 中间件将可以堆栈链中的任何位置将 panic 恢复，并将处理集中到
	app.Use(recover.New())
	// 设置路由
	routers.SetRoute(app)
	// 初始化 数据表
	initalize.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.FileInfo{})
	AutoErr := initalize.DB.AutoMigrate(models.NewFileInfo())
	if AutoErr != nil {
		fmt.Println("mysql迁移表失败")
	}
	// 监听端口
	_ = app.Listen(viper.GetString("App.Port"))
}

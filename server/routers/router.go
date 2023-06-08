package routers

import (
	v1 "fiber-layout/controllers/v1"
	"fiber-layout/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetRoute(app *fiber.App) {
	main := v1.NewDefaultController()
	group := app.Group("/v1")
	group.Post("/list", main.GetList)       // 获取文件列表
	group.Post("/get", main.GetFile)        // 获取文件信息
	group.Get("/download", main.Download)   // 下载文件
	group.Post("/login", main.Login)        // 登录
	group.Get("/randomImg", main.RandomImg) // 随机图片 需要传随机数 /v1/randomImg?12345678910

	// 以下接口需要权限
	group.Use(middleware.Auth)
	group.Post("/upload/chunkFile", main.ChunkFile)   // 上传切片
	group.Post("/upload/mergeFile", main.MergeFile)   // 合并切片
	group.Post("/upload/verifyFile", main.VerifyFile) // 检查file是否存在
}

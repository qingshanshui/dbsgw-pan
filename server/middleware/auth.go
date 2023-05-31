package middleware

import (
	"fiber-layout/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// Auth 权限晓验
func Auth(ctx *fiber.Ctx) error {
	token := ctx.Get("Authorization")
	if token == "" {
		return ctx.JSON(fiber.Map{
			"code": 0,
			"data": "token 不存在",
			"msg":  "操作失败",
		})
	} else {
		_, err := utils.ParseToken(token, viper.GetString("Jwt.Secret"))
		if err != nil {
			return ctx.JSON(fiber.Map{
				"code": 0,
				"data": "token 解析失败",
				"msg":  "操作失败",
			})
		} else {
			return ctx.Next()
		}
	}
}

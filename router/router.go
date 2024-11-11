package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"tiktok_e-commence/api"
	"tiktok_e-commence/middleware"
	"tiktok_e-commence/resp"
)

func InitRouter() {
	_ = godotenv.Load() // 加载环境变量.env
	h := server.Default(server.WithHostPorts(viper.GetString("server.port")))

	u := h.Group("/user")
	{
		u.POST("/register", api.Register)   // 注册
		u.POST("/login", api.Login)         // 登陆并获取双Token
		u.POST("refresh", api.RefreshToken) // 用refreshToken刷新双Token
		u.GET("/info", middleware.Auth, func(ctx context.Context, c *app.RequestContext) {
			res, _ := c.Get("ping")
			resp.Success(c, "ok", res.(string))
		}) // 测试鉴权中间件
	}

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.Spin()
}

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
)

func InitRouter() {
	_ = godotenv.Load() // 加载环境变量.env
	h := server.Default(server.WithHostPorts(viper.GetString("server.port")))

	u := h.Group("/user")
	{
		u.POST("/register", api.Register)
		u.POST("/login", api.Login)
	}

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.Spin()
}

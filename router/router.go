package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/spf13/viper"
)

func InitRouter() {
	port := viper.GetString("server.port")
	if port == "" {
		port = "8888"
	}
	h := server.Default(server.WithHostPorts(":" + port))

	//h.POST("/register", service.Register)          //注册
	//h.POST("/token", service.Login)                //登陆并获取双token
	//h.POST("token/refresh", service.RefreshToken)  //刷新token
	//h.GET("/info", middleWare.Token, service.Info) // 测试鉴权中间件，鉴权后可获取用户信息
	//h.Use(middleWare.Token)

	//test
	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.Spin()
}

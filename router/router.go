package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
	"tiktok_e-commence/api"
	"tiktok_e-commence/middleware"
)

func Router() {
	_ = godotenv.Load()

	h := server.Default(server.WithHostPorts(viper.GetString("server.port")))

	store, _ := redis.NewStore(10, "tcp", viper.GetString("redis.addr"), "", []byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("tiktok_e-commence", store))

	middleware.Register(h)

	h.LoadHTMLGlob("template/*")
	h.Static("static", "./")

	h.GET("/", api.Home)
	h.GET("/sign-in", func(ctx context.Context, c *app.RequestContext) { // 登陆
		data := utils.H{"Title": "Sign In", "Next": c.Query("next")}
		c.HTML(consts.StatusOK, "sign-in", data)
	})
	h.GET("/sign-up", func(ctx context.Context, c *app.RequestContext) { // 注册
		c.HTML(consts.StatusOK, "sign-up", utils.H{"Title": "Sign Up"})
	})
	h.GET("/about", func(ctx context.Context, c *app.RequestContext) {
		c.HTML(consts.StatusOK, "about", utils.H{"Title": "About"})
	})
	h.POST("/auth/login", api.Login)
	h.POST("/auth/register", api.Register)
	h.POST("/auth/logout", api.Logout)

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.Spin()
}

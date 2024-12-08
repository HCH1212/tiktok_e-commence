package router

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/monitor-prometheus"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"tiktok_e-commence/api"
	"tiktok_e-commence/common/mtl"
	"tiktok_e-commence/middleware"
	"tiktok_e-commence/resp"
)

func InitRouter() {
	// prometheus指标加到hertz
	r, registerInfo := mtl.InitMetric("hertz", ":9992", viper.GetString("consul.addr"))
	defer r.Deregister(registerInfo)

	// 链路追踪
	t := mtl.InitTracing("tiktok_e-commence")
	defer t.Shutdown(context.Background())
	tracer, cfg := hertztracing.NewServerTracer()

	_ = godotenv.Load() // 加载环境变量.env
	h := server.Default(server.WithHostPorts(viper.GetString("server.port")),
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
		tracer,
	)
	h.Use(hertztracing.ServerMiddleware(cfg))

	u := h.Group("/user")
	{
		u.POST("/register", api.Register)   // 注册
		u.POST("/login", api.Login)         // 登陆并获取双Token
		u.POST("refresh", api.RefreshToken) // 用refreshToken刷新双Token
		u.GET("/info", middleware.Auth, func(ctx context.Context, c *app.RequestContext) {
			res, _ := c.Get("id")
			resp.Success(c, "ok", res.(uint64))
		}) // 测试鉴权中间件
	}

	p := h.Group("/product")
	{
		p.POST("/add", api.CreateProduct)    // 添加商品
		p.POST("/update", api.ChangeProduct) // 更新商品
		p.POST("/delete", api.DeleteProduct) // 删除商品
		p.GET("/find", api.FindProduct)      // 精准查找某个商品bySUK
		p.GET("/search", api.FindProducts)   // 模糊搜索商品
	}

	c := h.Group("/cart", middleware.Auth)
	{
		c.POST("/add", api.AddItem)       // 添加商品到购物车
		c.POST("/delete", api.DeleteItem) // 删除购物车里的某个商品
		c.POST("/empty", api.EmptyCart)   // (物理)清空购物车
		c.GET("/get", api.GetCart)        // 查看购物车里的商品
	}

	o := h.Group("/order", middleware.Auth)
	{
		o.POST("/create", api.CreateOrder) // 创建订单
		o.GET("/list", api.ListOrder)      // 查看用户订单
		o.POST("/pay", api.IsPaidOrder)    // 设置订单已支付
	}

	pay := h.Group("/payment", middleware.Auth)
	{
		pay.POST("/charge", api.Charge) // 支付
	}

	h.GET("/ping", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(consts.StatusOK, utils.H{"message": "pong"})
	})
	h.Spin()
}

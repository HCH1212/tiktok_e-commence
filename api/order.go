package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"tiktok_e-commence/resp"
	"tiktok_e-commence/service"
)

func CreateOrder(ctx context.Context, c *app.RequestContext) {
	res, err := service.CreateOrderService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "创建订单成功", utils.H{"orderID": res.OrderId})
}

func ListOrder(ctx context.Context, c *app.RequestContext) {
	res, err := service.ListOrderService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "查询成功", res)
}

func IsPaidOrder(ctx context.Context, c *app.RequestContext) {
	err := service.IsPaidOrderService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "支付成功", nil)
}

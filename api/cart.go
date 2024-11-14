package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"tiktok_e-commence/resp"
	"tiktok_e-commence/service"
)

func AddItem(ctx context.Context, c *app.RequestContext) {
	err := service.AddItemService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "添加购物车成功", nil)
}

func DeleteItem(ctx context.Context, c *app.RequestContext) {
	err := service.DeleteItemService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "删除购物车商品成功", nil)
}

func EmptyCart(ctx context.Context, c *app.RequestContext) {
	err := service.EmptyCartService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "清空购物车成功", nil)
}

func GetCart(ctx context.Context, c *app.RequestContext) {
	res, err := service.GetCartService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "获取购物车商品成功", res)
}

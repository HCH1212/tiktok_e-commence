package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"tiktok_e-commence/resp"
	"tiktok_e-commence/service"
)

func CreateProduct(ctx context.Context, c *app.RequestContext) {
	res, err := service.CreateProductService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "创建成功", utils.H{"productId": res.Id})
}

func ChangeProduct(ctx context.Context, c *app.RequestContext) {
	res, err := service.ChangeProductService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "修改成功", utils.H{"productId": res.Id})
}

func DeleteProduct(ctx context.Context, c *app.RequestContext) {
	res, err := service.DeleteProductService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	if res {
		resp.Success(c, "删除成功", nil)
	} else {
		resp.Success(c, "删除失败", nil)
	}
}

func FindProduct(ctx context.Context, c *app.RequestContext) {
	res, err := service.FindProductService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "查找成功", res)
}

func FindProducts(ctx context.Context, c *app.RequestContext) {
	res, err := service.FindProductsService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "查找成功", res)
}

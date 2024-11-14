package service

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"tiktok_e-commence/model"
	"tiktok_e-commence/rpc"
)

func AddItemService(ctx context.Context, c *app.RequestContext) error {
	res, _ := c.Get("id")
	id := res.(uint64)
	suk := c.PostForm("suk")
	_, err := rpc.CartClient.AddItem(ctx, &cart.ItemReq{UserId: id, Suk: suk})
	if err != nil {
		return errors.New("rpc error")
	}
	return nil
}

func DeleteItemService(ctx context.Context, c *app.RequestContext) error {
	res, _ := c.Get("id")
	id := res.(uint64)
	suk := c.PostForm("suk")
	_, err := rpc.CartClient.DeleteItem(ctx, &cart.ItemReq{UserId: id, Suk: suk})
	if err != nil {
		return errors.New("rpc error")
	}
	return nil
}

func EmptyCartService(ctx context.Context, c *app.RequestContext) error {
	res, _ := c.Get("id")
	id := res.(uint64)
	_, err := rpc.CartClient.EmptyCart(ctx, &cart.UserId{UserId: id})
	if err != nil {
		return errors.New("rpc error")
	}
	return nil
}

func GetCartService(ctx context.Context, c *app.RequestContext) (resp []*model.ProductResp, err error) {
	res, _ := c.Get("id")
	id := res.(uint64)
	result, err := rpc.CartClient.GetCart(ctx, &cart.UserId{UserId: id})
	if err != nil {
		return nil, errors.New("rpc error")
	}
	resp = make([]*model.ProductResp, 0)
	for _, item := range result.Products {
		resp = append(resp, &model.ProductResp{
			SUK:         item.SUK,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			Price:       item.Price,
			Category:    item.Category,
		})
	}
	return
}

package service

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"tiktok_e-commence/model"
	"tiktok_e-commence/rpc"
)

func CreateOrderService(ctx context.Context, c *app.RequestContext) (resp *order.OrderId, err error) {
	id, _ := c.Get("id")
	userid := id.(uint64)
	suk := c.PostForm("suk")
	address := c.PostForm("address")
	res, err := rpc.OrderClient.CreateOrder(ctx, &order.Order{UserId: userid, Suk: suk, Address: address})
	if err != nil {
		return nil, errors.New("rpc error")
	}
	return res, nil
}

func ListOrderService(ctx context.Context, c *app.RequestContext) (resp []*model.Order, err error) {
	id, _ := c.Get("id")
	userid := id.(uint64)
	res, err := rpc.OrderClient.ListOrder(ctx, &order.UserId{UserId: userid})
	if err != nil {
		return nil, errors.New("rpc error")
	}
	resp = make([]*model.Order, 0)
	for _, item := range res.Orders {
		resp = append(resp, &model.Order{
			SUK:     item.Suk,
			Address: item.Address,
			IsPay:   item.IsPay,
		})
	}
	return
}

func IsPaidOrderService(ctx context.Context, c *app.RequestContext) error {
	orderId := c.PostForm("order_id")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		return errors.New("参数错误")
	}
	_, err = rpc.OrderClient.IsPaidOrder(ctx, &order.OrderId{OrderId: uint64(id)})
	if err != nil {
		return errors.New("rpc error")
	}
	return nil
}

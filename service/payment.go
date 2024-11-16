package service

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"tiktok_e-commence/rpc"
)

func ChargeService(ctx context.Context, c *app.RequestContext) (resp *payment.ChargeResp, err error) {
	id, _ := c.Get("id")
	userid := id.(uint64)
	amount, err := strconv.ParseFloat(c.PostForm("amount"), 32)
	if err != nil {
		return nil, errors.New("参数错误")
	}
	OrderId, err := strconv.Atoi(c.PostForm("order_id"))
	if err != nil {
		return nil, errors.New("参数错误")
	}
	cardNum := c.PostForm("card_num")
	res, err := rpc.PaymentClient.Charge(ctx, &payment.ChargeReq{Amount: float32(amount), UserId: userid, OrderId: uint64(OrderId), CardNum: cardNum})
	if err != nil {
		return nil, errors.New("rpc error")
	}
	return res, nil
}

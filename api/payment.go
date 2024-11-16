package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"tiktok_e-commence/resp"
	"tiktok_e-commence/service"
)

func Charge(ctx context.Context, c *app.RequestContext) {
	res, err := service.ChargeService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "支付成功", utils.H{"payment_id": res.PaymentId})
}

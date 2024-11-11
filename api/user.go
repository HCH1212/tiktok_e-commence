package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"tiktok_e-commence/resp"
	"tiktok_e-commence/service"
)

func Register(ctx context.Context, c *app.RequestContext) {
	err := service.RegisterService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
	}
	resp.Success(c, "register success", nil)
}

func Login(ctx context.Context, c *app.RequestContext) {
	res, err := service.LoginService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
	}
	resp.Success(c, "login success", utils.H{
		"accessToken":  res.AccessToken,
		"refreshToken": res.RefreshToken,
	})
}

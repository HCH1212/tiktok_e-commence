package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"tiktok_e-commence/resp"
	"tiktok_e-commence/service"
)

func Login(ctx context.Context, c *app.RequestContext) {
	res, err := service.LoginService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "login success", utils.H{
		"accessToken":  res.AccessToken,
		"refreshToken": res.RefreshToken,
	})
}

func Register(ctx context.Context, c *app.RequestContext) {
	err := service.RegisterService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "register success", nil)
}

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	res, err := service.RefreshService(ctx, c)
	if err != nil {
		resp.FailButServer(c, err.Error(), nil)
		return
	}
	resp.Success(c, "refresh token success", utils.H{
		"accessToken":  res.AccessToken,
		"refreshToken": res.RefreshToken,
	})
}

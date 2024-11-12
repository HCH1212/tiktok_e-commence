package service

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strings"
	"tiktok_e-commence/rpc"
)

func LoginService(ctx context.Context, c *app.RequestContext) (resp *auth.TwoToken, err error) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		return nil, errors.New("请求参数为空")
	}
	res, err := rpc.UserClient.Login(ctx, &user.LoginReq{Email: email, Password: password})
	if err != nil {
		log.Println("login:", err)
		return nil, err
	}
	// 获取双Token
	resp, err = rpc.AuthClient.GetToken(ctx, &auth.UserId{Id: res.Id})
	if err != nil {
		log.Println("get token:", err)
		return nil, err
	}
	return
}

func RegisterService(ctx context.Context, c *app.RequestContext) error {
	email := c.PostForm("email")
	password := c.PostForm("password")
	passwordAgain := c.PostForm("password_again")
	if email == "" || password == "" || passwordAgain == "" {
		return errors.New("请求参数为空")
	}
	_, err := rpc.UserClient.Register(ctx, &user.RegisterReq{Email: email, Password: password, PasswordAgain: passwordAgain})
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func RefreshService(ctx context.Context, c *app.RequestContext) (resp *auth.TwoToken, err error) {
	refreshToken := c.GetHeader("Authorization") // 传入refreshToken
	if len(refreshToken) == 0 || !strings.HasPrefix(string(refreshToken), "Bearer ") {
		return nil, errors.New("传入参数有误")
	}
	refreshToken = refreshToken[7:]

	resp, err = rpc.AuthClient.ExecRefreshToken(ctx, &auth.RefreshToken{RefreshToken: string(refreshToken)})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return
}

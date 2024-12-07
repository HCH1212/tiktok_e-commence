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
		if err.Error() == "remote or network error[remote]: biz error: 用户不存在" {
			return nil, errors.New("用户不存在")
		}
		if err.Error() == "remote or network error[remote]: biz error: 密码错误" {
			return nil, errors.New("密码错误")
		}
		return nil, errors.New("rpc error")
	}
	// 获取双Token
	resp, err = rpc.AuthClient.GetToken(ctx, &auth.UserId{Id: res.Id})
	if err != nil {
		log.Println("get token:", err)
		return nil, errors.New("rpc error")
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
	if password != passwordAgain {
		return errors.New("请求参数错误")
	}
	_, err := rpc.UserClient.Register(ctx, &user.RegisterReq{Email: email, Password: password, PasswordAgain: passwordAgain})
	if err != nil {
		log.Println(err)
		if err.Error() == "remote or network error: rpc error: code = 13 desc = 用户已存在 [biz error]" {
			return errors.New("用户已存在")
		}
		return errors.New("rpc error")
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
		return nil, errors.New("rpc error")
	}
	return
}

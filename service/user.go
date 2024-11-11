package service

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"tiktok_e-commence/rpc"
)

func LoginService(ctx context.Context, c *app.RequestContext) (resp *auth.TwoToken, err error) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		log.Println("请求参数为空")
		return nil, errors.New("请求参数为空")
	}
	res, err := rpc.UserClient.Login(ctx, &user.LoginReq{Email: email, Password: password})
	if err != nil {
		log.Println("login:", err)
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
		log.Println("请求参数为空")
		return errors.New("请求参数为空")
	}
	_, err := rpc.UserClient.Register(ctx, &user.RegisterReq{Email: email, Password: password, PasswordAgain: passwordAgain})
	if err != nil {
		log.Println(err)
		return errors.New("rpc error")
	}
	return nil
}

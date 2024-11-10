package service

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	"log"
	"tiktok_e-commence/rpc"
)

func LoginService(ctx context.Context, c *app.RequestContext) (redirect string, err error) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if email == "" || password == "" {
		log.Println("请求参数为空")
		return "", errors.New("请求参数为空")
	}
	_, err = rpc.UserClient.Login(ctx, &user.LoginReq{Email: email, Password: password})
	if err != nil {
		log.Println(err)
		return "", errors.New("rpc error")
	}

	next := c.Query("next")
	redirect = "/"
	if next != "" {
		redirect = next
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

func LogoutService(c *app.RequestContext) (string, error) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		return "", err
	}
	return "???????????????????????", nil
}

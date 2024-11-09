package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"tiktok_e-commence/model"
	"tiktok_e-commence/service"
)

func Login(ctx context.Context, c *app.RequestContext) {
	var user model.User
	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	if user.Email == "" || user.Password == "" {
		log.Println("empty")
		c.JSON(consts.StatusBadRequest, "empty")
		return
	}
	err := c.BindAndValidate(&user)
	if err != nil {
		log.Println(err)
		c.JSON(consts.StatusBadRequest, "error")
		return
	}
	redirect, err := service.LoginService(c)
	if err != nil {
		log.Println(err)
		c.JSON(consts.StatusServiceUnavailable, "error")
		return
	}
	c.Redirect(consts.StatusOK, []byte(redirect))
}

func Register(ctx context.Context, c *app.RequestContext) {
	var userRe model.UserRe
	userRe.Email = c.PostForm("email")
	userRe.Password = c.PostForm("password")
	userRe.PasswordConfirm = c.PostForm("password_confirm")
	if userRe.Email == "" || userRe.Password == "" || userRe.PasswordConfirm == "" {
		log.Println("empty")
		c.JSON(consts.StatusBadRequest, "empty")
		return
	}
	err := c.BindAndValidate(&userRe)
	if err != nil {
		log.Println(err)
		c.JSON(consts.StatusBadRequest, "error")
		return
	}
	_, err = service.RegisterService(c)
	if err != nil {
		log.Println(err)
		c.JSON(consts.StatusServiceUnavailable, "error")
		return
	}
	c.Redirect(consts.StatusOK, []byte("/"))
}

func Logout(ctx context.Context, c *app.RequestContext) {
	_, err := service.LogoutService(c)
	if err != nil {
		log.Println(err)
		c.JSON(consts.StatusServiceUnavailable, "error")
		return
	}
	c.Redirect(consts.StatusOK, []byte("/"))
}

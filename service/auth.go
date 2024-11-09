package service

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func LoginService(c *app.RequestContext) (redirect string, err error) {
	session := sessions.Default(c)
	session.Set("user_id", 1)
	err = session.Save()
	if err != nil {
		return "", err
	}
	next := c.Query("next")
	redirect = "/"
	if next != "" {
		redirect = next
	}
	return
}

func RegisterService(c *app.RequestContext) (string, error) {
	session := sessions.Default(c)
	session.Set("user_id", 1)
	err := session.Save()
	if err != nil {
		return "", err
	}
	return "???????????????????????", nil
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

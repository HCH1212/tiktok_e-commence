package api

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"tiktok_e-commence/service"
	"tiktok_e-commence/utils"
)

func Home(ctx context.Context, c *app.RequestContext) {
	res, err := service.HomeService(c)
	if err != nil {
		log.Println(err)
		return
	}
	c.HTML(consts.StatusOK, "home", utils.AddUserId(c, res))
}

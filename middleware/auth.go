package middleware

import (
	"context"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
	"tiktok_e-commence/resp"
	rpc "tiktok_e-commence/rpc"
)

func Auth(ctx context.Context, c *app.RequestContext) {
	token := c.GetHeader("Authorization") // 传入accessToken

	if len(token) == 0 || !strings.HasPrefix(string(token), "Bearer ") {
		resp.Fail(c, "传入参数有误", nil)
		return
	}

	token = token[7:]

	res, err := rpc.AuthClient.VerifyToken(ctx, &auth.AccessToken{AccessToken: string(token)})
	if err != nil {
		resp.FailButServer(c, "权限不足", nil)
		return
	}

	// 鉴权成功, 将用户id设置为上下文
	c.Set("id", res.Id)
	c.Next(ctx)
}

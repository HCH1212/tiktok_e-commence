package utils

import (
	"github.com/cloudwego/hertz/pkg/app"
)

// 自动添加user_id
func AddUserId(c *app.RequestContext, content map[string]any) map[string]any {
	content["user_id"], _ = c.Get("user_id")
	return content
}

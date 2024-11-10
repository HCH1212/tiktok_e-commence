package resp

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

const (
	Ok         = 2000
	No         = 4000
	ServerFail = 5000
)

// OkWithData 成功的返回
func Success(c *app.RequestContext, message string, data interface{}) {
	ResponseWithStatusAndData(c, Ok, message, data)
	c.Abort()
}

// FailWithData 客户端请求失败的返回
func Fail(c *app.RequestContext, message string, data interface{}) {
	ResponseWithStatusAndData(c, No, message, data)
	c.Abort()
}

// ServerFailWithData 服务端响应失败的返回
func FailButServer(c *app.RequestContext, message string, data interface{}) {
	ResponseWithStatusAndData(c, ServerFail, message, data)
	c.Abort()
}

// ResponseWithStatusAndData 确定统一返回格式
func ResponseWithStatusAndData(c *app.RequestContext, status int, message string, data interface{}) {
	c.JSON(200, utils.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

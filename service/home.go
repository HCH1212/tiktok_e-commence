package service

import "github.com/cloudwego/hertz/pkg/app"

func HomeService(c *app.RequestContext) (map[string]any, error) {
	res := make(map[string]any)
	items := []map[string]any{
		{"Name": "Rose", "Price": 120, "Picture": "static/image/R-C.jpeg"},
		{"Name": "Beautiful girl", "Price": 100, "Picture": "static/image/xun.jpg"},
	}
	res["Title"] = "Hot Sales"
	res["Items"] = items
	return res, nil
}

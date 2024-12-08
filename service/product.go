package service

import (
	"context"
	"errors"
	"github.com/HCH1212/tiktok_e-commence_rpc/gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/sirupsen/logrus"
	"strconv"
	"tiktok_e-commence/model"
	"tiktok_e-commence/rpc"
)

func CreateProductService(ctx context.Context, c *app.RequestContext) (resp *product.ProductId, err error) {
	var pro model.Product
	err = c.BindJSON(&pro)
	if err != nil {
		return nil, errors.New("参数错误")
	}
	res, err := rpc.ProductClient.CreateProduct(ctx, &product.Product{
		SUK:         pro.SUK,
		Name:        pro.Name,
		Description: pro.Description,
		Picture:     pro.Picture,
		Price:       pro.Price,
		Category:    pro.Category,
	})
	if err != nil {
		logrus.Println(err)
		if err.Error() == "remote or network error[remote]: biz error: 商品已存在" {
			return nil, errors.New("商品已存在")
		}
		return nil, errors.New("rpc error")
	}
	resp = &product.ProductId{Id: res.Id}
	return
}

func ChangeProductService(ctx context.Context, c *app.RequestContext) (resp *product.ProductId, err error) {
	var pro model.Product
	err = c.BindJSON(&pro)
	if err != nil {
		return nil, errors.New("参数错误")
	}
	res, err := rpc.ProductClient.ChangeProduct(ctx, &product.Product{
		SUK:         pro.SUK,
		Name:        pro.Name,
		Description: pro.Description,
		Picture:     pro.Picture,
		Price:       pro.Price,
		Category:    pro.Category,
	})
	if err != nil {
		logrus.Println(err)
		if err.Error() == "remote or network error[remote]: biz error: 商品不存在" {
			return nil, errors.New("商品不存在")
		}
		return nil, errors.New("rpc error")
	}
	resp = &product.ProductId{Id: res.Id}
	return
}

func DeleteProductService(ctx context.Context, c *app.RequestContext) (bool, error) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	res, err := rpc.ProductClient.DeleteProduct(ctx, &product.ProductId{Id: uint64(id)})
	if err != nil || !res.Pass {
		logrus.Println(err)
		return false, errors.New("rpc error")
	}
	return true, nil
}

func FindProductService(ctx context.Context, c *app.RequestContext) (resp *model.ProductResp, err error) {
	suk := c.PostForm("suk")
	res, err := rpc.ProductClient.FindProduct(ctx, &product.ProductSUK{SUK: suk})
	if err != nil {
		logrus.Println(err)
		if err.Error() == "remote or network error[remote]: biz error: 商品不存在" {
			return nil, errors.New("商品不存在")
		}
		return nil, errors.New("rpc error")
	}
	resp = &model.ProductResp{
		SUK:         res.SUK,
		Name:        res.Name,
		Description: res.Description,
		Picture:     res.Picture,
		Price:       res.Price,
		Category:    res.Category,
	}
	return
}

func FindProductsService(ctx context.Context, c *app.RequestContext) (resp []*model.ProductResp, err error) {
	name := c.PostForm("name")
	res, err := rpc.ProductClient.FindProducts(ctx, &product.SearchReq{Name: name})
	if err != nil {
		logrus.Println(err)
		if err.Error() == "remote or network error[remote]: biz error: 商品不存在" {
			return nil, errors.New("商品不存在")
		}
		return nil, errors.New("rpc error")
	}
	resp = make([]*model.ProductResp, 0)
	for _, item := range res.Products {
		resp = append(resp, &model.ProductResp{
			SUK:         item.SUK,
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			Price:       item.Price,
			Category:    item.Category,
		})
	}
	return
}

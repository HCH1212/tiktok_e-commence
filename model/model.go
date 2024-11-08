package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string // 用户名
	Password    string // 密码
	ShoppingCar        // 购物车
	// 还有一个用户的权限认证
}

type Goods struct {
	gorm.Model
	Name        string // 商品名
	Description string // 描述
	Price       int    // 价格
	Amount      int    // 库存
}

type ShoppingCar struct {
	gorm.Model
	MyGoods []Goods // 商品
}

type Record struct {
	gorm.Model
	Goods // 订单的商品
	User  // 订单的用户
	// 订单的过期可以用redis来处理,订单感觉可以不用存入数据库
}

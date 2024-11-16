package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Product struct {
	gorm.Model
	SUK         string   `json:"suk"`
	Name        string   `json:"name"`
	Price       float32  `json:"price"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Category    []string `json:"category"`
}

type ProductResp struct {
	SUK         string   `json:"suk"`
	Name        string   `json:"name"`
	Price       float32  `json:"price"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Category    []string `json:"category"`
}

type Order struct {
	SUK     string `json:"suk"`
	Address string `json:"address"`
	IsPay   bool   `json:"is_pay"`
}

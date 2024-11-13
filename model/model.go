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
	Price       int64    `json:"price"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Category    []string `json:"category"`
}

type ProductResp struct {
	SUK         string   `json:"suk"`
	Name        string   `json:"name"`
	Price       int64    `json:"price"`
	Description string   `json:"description"`
	Picture     string   `json:"picture"`
	Category    []string `json:"category"`
}

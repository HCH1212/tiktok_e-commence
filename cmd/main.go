package main

import (
	"tiktok_e-commence/config"
	"tiktok_e-commence/router"
)

func main() {
	config.InitConfig()
	router.Router()
}

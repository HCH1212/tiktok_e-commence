package main

import (
	"tiktok_e-commence/config"
	"tiktok_e-commence/router"
	"tiktok_e-commence/rpc"
)

func main() {
	config.InitConfig()
	rpc.InitClient()
	router.InitRouter()
}

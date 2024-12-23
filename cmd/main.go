package main

import (
	"tiktok_e-commence/config"
	"tiktok_e-commence/log"
	"tiktok_e-commence/router"
	"tiktok_e-commence/rpc"
)

func main() {
	config.InitConfig()
	// 自定义log
	log.InitDefaultLogger()
	rpc.InitClient()
	router.InitRouter()
}

package main

import (
	"4-TikTok/IM"
	"4-TikTok/configs"
	"4-TikTok/router"
)

func main() {
	configs.Init()
	go IM.Manager.Start()
	r := router.NewRouter()
	_ = r.Run(":10001")
}

package main

import (
	"4-TikTok/IM"
	"4-TikTok/configs"
	"4-TikTok/router"
	"fmt"
)

func main() {
	configs.Init()
	go IM.Manager.Start()
	r := router.NewRouter()
	fmt.Println("这是一次PR尝试")
	_ = r.Run(":10001")
}

package handler

import (
	"4-TikTok/service"
	utils "4-TikTok/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetVideoFeed(c *gin.Context) {
	var Feed service.FeedService
	if err := c.ShouldBind(&Feed); err == nil {
		res := Feed.GetVideoListByFeed()
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

func VideoFeedList(c *gin.Context) {
	var List service.VideoFeedList

	if err := c.ShouldBind(&List); err == nil {
		res := List.GetVideoFeedList()
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

func VideoSearch(c *gin.Context) {
	var Search service.VideoSearch
	claims, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&Search); err == nil {
		res := Search.VideoSearch(claims.Id)
		c.JSON(200, res)
	} else {
		fmt.Println("出错了")
		c.JSON(400, "err")
	}
}

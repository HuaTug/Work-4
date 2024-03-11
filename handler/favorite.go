package handler

import (
	"4-TikTok/service"
	utils "4-TikTok/util"
	"github.com/gin-gonic/gin"
)

func Favorite(c *gin.Context) {
	var favorite service.FavoriteService
	claims, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&favorite); err == nil {
		res := favorite.Favorite(claims.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

func FavorList(c *gin.Context) {
	var favorite service.FavoriteListService
	claims, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&favorite); err == nil {
		res := favorite.List(claims.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

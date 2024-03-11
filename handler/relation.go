package handler

import (
	"4-TikTok/service"
	utils "4-TikTok/util"
	"github.com/gin-gonic/gin"
)

func Following(c *gin.Context) {
	var relation service.FollowService
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&relation); err == nil {
		res := relation.Following(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}
func List(c *gin.Context) {
	var relation service.FollowServicePage
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&relation); err == nil {
		res := relation.List(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}
func List2(c *gin.Context) {
	var relation service.FollowServicePage
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&relation); err == nil {
		res := relation.List2(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

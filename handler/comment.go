package handler

import (
	"4-TikTok/service"
	utils "4-TikTok/util"
	"github.com/gin-gonic/gin"
)

func CommentCreate(c *gin.Context) {
	var comment service.CommentService
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&comment); err == nil {
		res := comment.Create(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

func CommentDelete(c *gin.Context) {
	var comment service.CommentService
	if err := c.ShouldBind(&comment); err == nil {
		res := comment.Delete()
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

func CommentList(c *gin.Context) {
	var list service.ListCommentsService
	claims, _ := utils.ParseToken(c.GetHeader("token"))
	if err := c.ShouldBind(&list); err == nil {
		res := list.List(claims.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, "err")
	}
}

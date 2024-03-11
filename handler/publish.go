package handler

import (
	"4-TikTok/service"
	utils "4-TikTok/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	var upload service.UploadVideoService
	claim, _ := utils.ParseToken(c.GetHeader("token"))
	if file, err := c.FormFile("file"); err != nil {
		fmt.Println("上传file文件出错")
		panic(err)
	} else {
		if err := c.ShouldBind(&upload); err == nil {
			res := upload.UploadVideo(file, claim.Id)
			c.JSON(200, res)
		} else {
			c.JSON(500, "err")
		}
	}
}

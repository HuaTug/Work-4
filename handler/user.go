package handler

import (
	"4-TikTok/service"
	utils "4-TikTok/util"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

func UserRegister(c *gin.Context) {
	var user service.UserService
	if err := c.ShouldBind(&user); err == nil {
		res := user.Register()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(400, "err")
	}
}

func UserLogin(c *gin.Context) {
	var userLoginService service.UserService
	if err := c.ShouldBind(&userLoginService); err == nil {
		res := userLoginService.Login()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(400, "err")
	}
}

func UserInfo(c *gin.Context) {
	var info service.InfoService
	if err := c.ShouldBind(&info); err == nil {
		res := info.Info()
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(400, "err")
	}
}

func UpLoad(c *gin.Context) {
	claim, err := utils.ParseToken(c.GetHeader("token")) //要合理使用token，token包含着登录用户的信息
	formFile, err := c.FormFile("file")
	if err != nil {
		logging.Info(err)
		c.JSON(500, "err")
	}
	if err == nil {
		res := service.UploadImage(formFile, claim.UserName)
		c.JSON(200, res)
	} else {
		logging.Info(err)
		c.JSON(400, "err")
	}
}

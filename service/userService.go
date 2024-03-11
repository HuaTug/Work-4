package service

import (
	"4-TikTok/model"
	"4-TikTok/serializer"
	utils "4-TikTok/util"
	"fmt"
	"github.com/jinzhu/gorm"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

type UserService struct {
	UserName string `form:"username" json:"username" binding:"required,min=4,max=10"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=12"`
}

type InfoService struct {
	Uid      int64  `form:"uid" json:"uid"`
	UserName string `form:"username" json:"username"`
}

func (service *UserService) Register() serializer.Response {
	var user model.User
	var count int64
	model.DB.Model(&model.User{}).Where("username=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return serializer.Response{
			Status: 400,
			Msg:    "已经有人了，无需再注册",
		}
	}
	user.UserName = service.UserName
	err := user.SetPassword(service.Password)
	if err != nil {
		return serializer.Response{}
	}
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "数据库操作错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "用户注册成功",
	}
}

func (service *UserService) Login() serializer.Response {
	var user model.User
	fmt.Println(service.UserName)
	if err := model.DB.Where("username=?", service.UserName).First(&user).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return serializer.Response{
				Status: 400,
				Msg:    "用户不存在，请先注册",
			}
		}
		return serializer.Response{
			Status: 500,
			Msg:    "数据库错误",
		}
	}
	if !user.CheckPassword(service.Password) {
		return serializer.Response{
			Status: 400,
			Msg:    "密码错误",
		}
	}
	token, err := utils.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "Token签发错误",
		}
	}
	return serializer.Response{
		Status: 200,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    "登录成功",
	}
}

func (info *InfoService) Info() serializer.Response {
	var user model.User
	code := 200
	fmt.Println(info.UserName)
	err := model.DB.Where("username=?", info.UserName).First(&user).Error
	if err != nil {
		code = 500
		return serializer.Response{
			Status: code,
			Msg:    "查询失败",
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildUser(user),
	}
}

func UploadImage(file *multipart.FileHeader, username string) serializer.Response {
	var user model.User
	var info UserService
	src, err := file.Open()
	if err != nil {
		fmt.Println("第一步出现了错误")
		return serializer.Response{}
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
		}
	}(src)

	fileName := time.Now().Unix()
	uploadDir := "./ImageUpload"
	dstPath := fmt.Sprintf("%s/%s.jpg", uploadDir, strconv.FormatInt(fileName, 10)) //ToDo 注意这里需要一个完整的路径地址，且该路径需要存在
	dst, err := os.Create(dstPath)
	if err != nil {
		fmt.Println("第二步出现了错误")
		return serializer.Response{}

	}
	defer func(dst *os.File) {
		err := dst.Close()
		if err != nil {
		}
	}(dst)

	_, err = io.Copy(dst, src)
	if err != nil {
		fmt.Println("第三步出现了错误")
		return serializer.Response{}
	}

	imageURL := fmt.Sprintf("https://www.Nwu.com/uploads/%s", fileName)
	user.Avatar = imageURL
	fmt.Println(info.UserName)
	//ToDo: 这是更新的GORM操作
	err = model.DB.Model(model.User{}).Where("username=?", username).Update(model.User{
		Avatar: imageURL,
	}).Error
	if err != nil {
		panic(err)
	} else {
		return serializer.Response{
			Status: 200,
			Data:   serializer.BuildUser(user),
		}
	}
}

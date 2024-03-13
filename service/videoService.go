package service

import (
	"4-TikTok/model"
	"4-TikTok/serializer"
	"fmt"
	logging "github.com/sirupsen/logrus"
)

type FeedService struct {
	LastTime string `form:"last_time" json:"last_time"`
}

type VideoFeedList struct {
	AuthorId uint  `form:"author_id" json:"author_id"`
	PageNum  int64 `form:"page_num" json:"page_num"`
	PageSize int64 `form:"page_size" json:"page_size"`
}

type VideoSearch struct {
	KeyWords string `form:"key_words" json:"key_words" binding:"required"`
	PageSize int64  `form:"page_size" json:"page_size" binding:"required"`
	PageNum  int64  `form:"page_num" json:"page_num" binding:"required"`
	FromDate int64  `form:"from_date" json:"from_date" `
	ToDate   int64  `form:"to_date" json:"to_date" `
	UserName string `form:"username" json:"username"`
}

type VideoPopular struct {
	PageNum  int64 `form:"page_num" json:"page_num"`
	PageSize int64 `form:"page_size" json:"page_size"`
}

func (service *FeedService) GetVideoListByFeed() serializer.Response {
	var Videos []model.Video
	var count int64
	fmt.Println(service.LastTime)
	if err := model.DB.Model(&Videos).Where("publish_time<?", service.LastTime).Limit(20).Order("video_id DESC").Count(&count).Find(&Videos).Error; err != nil {
		return serializer.Response{
			Status: 500,
			Msg:    "err",
		}
	}
	return serializer.BuildListResponse(Videos, uint(count))
}

func (service *VideoFeedList) GetVideoFeedList() serializer.Response {
	var Videos []model.Video
	var count int64
	fmt.Println(service.AuthorId, service.PageSize, service.PageNum) //Todo 用于检测哪里出现错误
	//ToDo :在分页操作中遇到了第二页无法查询的问题 修正:count语句查询不能再分页查询limit和offset之后
	if err := model.DB.Model(&model.Video{}).Preload("Author").Where("author_id=?", service.AuthorId).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).
		Find(&Videos).Error; err != nil {
		logging.Error(err)
		return serializer.Response{
			Status: 500,
			Msg:    "err",
		}
	}
	return serializer.BuildListResponse(Videos, uint(count))
}

func (service *VideoSearch) VideoSearch(uid uint) serializer.Response {
	var Videos []model.Video
	var view model.Video
	var count uint
	if err := model.DB.Model(&model.Video{}).Where("title like ?", "%"+service.KeyWords+"%").
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Count(&count).Find(&Videos).Error; err != nil {
		fmt.Println("搜索发生了问题")
		return serializer.Response{
			Status: 500,
			Msg:    "err",
		}
	}
	view.Addview(uid)
	return serializer.BuildListResponse(Videos, uint(count))
}

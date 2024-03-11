package service

import (
	"4-TikTok/model"
	"4-TikTok/serializer"
	"fmt"
)

type FavoriteService struct {
	VideoID    int64 `form:"video_id" json:"video_id"`
	CommentID  int64 `form:"comment_id" json:"comment_id"`
	ActionType int   `form:"action_type" json:"action_type"`
}

type FavoriteListService struct {
}

func (service *FavoriteService) Favorite(uid uint) serializer.Response {
	var favorite model.Favorite
	favorites := &model.Favorite{
		VideoId:    service.VideoID,
		CommentId:  service.CommentID,
		ActionType: service.ActionType,
		UserId:     uid,
	}
	if err := model.DB.Model(&favorite).Create(favorites).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Msg:    "创建失败!",
		}
	}
	//ToDo 如何修改为不同的用户可以点赞或者取消
	if err := model.DB.Model(&model.Video{}).Where("video_id=? and author_id=?", service.VideoID, uid).
		Update("favorite_count", favorite.Video.FavoriteCount+1).Error; err != nil { //这里是对数据库Update的操作
		fmt.Println("对视频的点赞失败")
	}
	fmt.Println("对视频的点赞成功")
	return serializer.Response{
		Status: 200,
		Msg:    "创建成功",
	}
}

func (service *FavoriteListService) List(uid uint) serializer.Response {
	var fav []model.Favorite
	var total uint
	model.DB.Model(&model.Favorite{}).Preload("User").Where("user_id=?", uid).Count(&total).Find(&fav)
	return serializer.BuildListResponse(fav, total)
}

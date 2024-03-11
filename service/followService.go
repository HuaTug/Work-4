package service

import (
	"4-TikTok/model"
	"4-TikTok/serializer"
	"fmt"
)

type FollowService struct {
	TouserId   int `form:"to_user_id" json:"to_user_id"`
	ActionType int `form:"action_type" json:"action_type"`
}
type FollowServicePage struct {
	UserId   int   `form:"user_id" json:"user_id"`
	PageNum  int64 `form:"page_num" json:"page_num"`
	PageSize int64 `form:"page_size" json:"page_size"`
}

func (service *FollowService) Following(uid uint) serializer.Response {
	var Relationservice model.Relation
	Relation := &model.Relation{
		Follower:    service.TouserId,
		Action_type: service.ActionType,
		UserId:      int(uid),
	}
	if service.ActionType == 1 {
		if err := model.DB.Model(&model.Relation{}).Create(Relation).Error; err != nil {
			fmt.Println("关注操作出错")
			return serializer.Response{
				Status: 500,
				Msg:    "关注操作出错",
			}
		}
	} else {
		model.DB.Model(&Relationservice).Where("follower_id=?", service.TouserId).Find(&Relationservice)
		Relationservice.Action_type = service.ActionType
		if err := model.DB.Save(&Relationservice).Error; err != nil {
			fmt.Println("取消关注出错")
			return serializer.Response{
				Status: 500,
				Msg:    "取消关注出错",
			}
		}
		return serializer.Response{
			Status: 200,
			Msg:    "取消关注成功",
		}
	}
	if err := model.DB.Model(&model.Relation{}).Where("user_id=?", service.TouserId).Find(&Relationservice).Error; err != nil {
		fmt.Println("未找到这个人")
	}
	Relationservice.Follow = int(uid)
	model.DB.Save(&Relationservice)
	return serializer.Response{
		Status: 200,
		Msg:    "关注成功",
	}
}

func (service *FollowServicePage) List(uid uint) serializer.Response {
	var list []model.Relation
	var count uint
	if err := model.DB.Model(&model.Relation{}).Preload("User").Where("follower_id=?", uid).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Count(&count).Find(&list).Error; err != nil {
		fmt.Println("分页查询出错")
		return serializer.Response{
			Status: 500,
			Msg:    "分页查询出错",
		}
	}
	return serializer.BuildListResponse(list, count)
}

func (service *FollowServicePage) List2(uid uint) serializer.Response {
	var list []model.Relation
	var count uint
	if err := model.DB.Model(&model.Relation{}).Preload("User").Where("follow_id=?", uid).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Count(&count).Find(&list).Error; err != nil {
		fmt.Println("分页查询出错")
		return serializer.Response{
			Status: 500,
			Msg:    "分页查询出错",
		}
	}
	return serializer.BuildListResponse(list, count)
}

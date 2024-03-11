package service

import (
	"4-TikTok/model"
	"4-TikTok/serializer"
)

type CommentService struct {
	VideoId   int64  `form:"video_id" json:"video_id"`
	CommentId int64  `form:"comment_id" json:"comment_id"`
	Comment   string `form:"comment" json:"comment"`
}

type ListCommentsService struct {
	PageNum  int `json:"page_num" form:"page_num"`
	PageSize int `json:"page_size" form:"page_size"`
}

func (service *CommentService) Create(uid uint) serializer.Response {
	comment := &model.Comment{
		VideoId:   service.VideoId,
		CommentId: service.CommentId,
		Comment:   service.Comment,
		UserId:    uid,
	}
	err := model.DB.Model(&model.Comment{}).Create(comment).Error
	if err != nil {
		return serializer.Response{
			Status: 400,
			Error:  "err",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "成功创建一条留言",
	}
}

func (service *CommentService) Delete() serializer.Response {
	var comment model.Comment
	if err := model.DB.Delete(&comment, "video_id=? And comment_id=?", service.VideoId, service.CommentId).Error; err != nil {
		return serializer.Response{
			Status: 400,
			Error:  "err",
		}
	}
	return serializer.Response{
		Status: 200,
		Msg:    "成功删除一条记录",
	}
}

func (service *ListCommentsService) List(uid uint) serializer.Response {
	var comments []model.Comment
	var total int64
	if service.PageSize == 0 {
		service.PageSize = 10
	}
	model.DB.Model(&model.Comment{}).Preload("User").Where("user_id=?", uid).Count(&total).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&comments)
	return serializer.BuildListResponse(comments, uint(total))
}

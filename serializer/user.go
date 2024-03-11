package serializer

import "4-TikTok/model"

type User struct {
	ID        uint   `json:"id" form:"id" example:"1"`
	UserName  string `json:"user_name" form:"user_name" example:"NigTusg"`
	CreateAt  int64  `json:"create_at" form:"create_at"`
	UpdatedAt int64  `json:"updated_at" form:"update_at"`
	DeletedAt int64  `json:"deleted_at" form:"deleted_at"`
	Avatar    string `json:"avatar_url" form:"avatar_url"`
}

func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
		Avatar:   user.Avatar,
	}
}

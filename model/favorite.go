package model

type Favorite struct {
	Id         uint `gorm:"column:favorite_id;primary_key"`
	UserId     uint
	VideoId    int64
	CommentId  int64
	ActionType int
	User       User  `gorm:"Foreign_key:UserId"`
	Video      Video `gorm:"foreign_key:VideoId"`
}

package model

import "github.com/jinzhu/gorm"

type Comment struct {
	VideoId   int64
	CommentId int64
	UserId    uint
	Comment   string
	LikeCount string
	Time      string
	User      User `gorm:"Foreignkey:UserId"`
	gorm.Model
}

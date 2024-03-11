package model

type Relation struct {
	Id          int  `gorm:"primary_key"`
	UserId      int  `gorm:"user_id"`
	Follow      int  `gorm:"column:follow_id"`
	Follower    int  `gorm:"column:follower_id"`
	Action_type int  `gorm:"column:action_type"`
	User        User `gorm:"foreign_key:UserId"`
}

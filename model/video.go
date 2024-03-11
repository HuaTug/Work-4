package model

import (
	"4-TikTok/cache"
	"strconv"
	"time"
)

type Video struct {
	Id            int64 `gorm:"column:video_id;primary_key"`
	AuthorId      uint  `gorm:"column:author_id"`
	PlayUrl       string
	CoverUrl      string
	FavoriteCount int64
	CommentCount  int64
	PublishTime   time.Time
	Title         string
	Author        User `gorm:"foreign_key:AuthorId"` //ToDo: 需要满足关联的外键类型和被关联模型User的主键类型相同，即为uint
}

// 使用redis缓存完成对于视频的访问次数
// ToDo 后续还需要再继续对其完善

func (Task *Video) Addview(uid uint) {
	cache.RedisClient.Incr(cache.TaskViewKey(uid))
	cache.RedisClient.ZIncrBy(cache.Rankey, 1, strconv.Itoa(int(Task.Id)))
}
